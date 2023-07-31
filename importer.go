package gormy

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type _importer struct {
	schemaName     string
	ignoredTables  []string
	importedTables []importedTable
	columnNameMap  func(string) string
	schemaMap      map[string]string
	outputDest     string
}

func (importer *_importer) Import() error {
	err := importer.DiscoverTables()

	if err != nil {
		return fmt.Errorf("discovering tables: %w", err)
	}

	for _, it := range importer.importedTables {
		err = it.ToFile(importer.outputDest)
		if err != nil {
			return fmt.Errorf("discovering tables: %w", err)
		}
	}

	return nil
}

func (importer *_importer) DiscoverTables() error {

	itables := []importedTable{}
	tableNames := []string{}

	rows, err := gc.conn.Query(
		fmt.Sprintf(`SELECT table_name FROM information_schema.tables
		WHERE table_schema='%s'`,
			importer.schemaName),
	)

	if err != nil {
		return fmt.Errorf("querying database schema: %w", err)
	}

	for {
		ok := rows.Next()

		if !ok {
			break
		}

		row := new(sql.NullString)

		err = rows.Scan(&row)

		if err != nil {
			return fmt.Errorf("discovering tables: %w", err)
		}

		ignored := false

		for _, v := range importer.ignoredTables {
			if v == row.String {
				ignored = true
			}
		}

		if !ignored {
			tableNames = append(tableNames, row.String)
		}
	}

	for _, tableName := range tableNames {
		icols, err := importer.discoverColumns(tableName)

		if err != nil {
			return err
		}

		itable := importedTable{
			tableName:  tableName,
			structName: DefaultNameMap(tableName),
			columns:    icols,
		}

		itables = append(itables, itable)

	}

	importer.importedTables = itables

	return nil
}

func (importer *_importer) discoverColumns(tableName string) ([]importedColumn, error) {
	rows, err := gc.conn.Query(
		fmt.Sprintf(`SELECT column_name as sql_name, udt_name as sql_type FROM information_schema.columns
		WHERE table_schema='public' and table_name = '%s'`,
			tableName),
	)

	if err != nil {
		return nil, fmt.Errorf("querying table schema: %w", err)
	}

	columnTypes, err := rows.ColumnTypes()
	columns := []map[string]string{}

	if err != nil {
		return nil, fmt.Errorf("reading column types: %w", err)
	}

	for {
		ok := rows.Next()

		if !ok {
			break
		}

		row := make([]interface{}, 2)

		for i, _ := range row {
			row[i] = new(sql.NullString)
		}

		err = rows.Scan(row...)

		if err != nil {
			return nil, fmt.Errorf("scanning column data: %w", err)
		}

		column := map[string]string{}

		for i, col := range columnTypes {
			if z, isStr := row[i].(*sql.NullString); isStr {
				column[col.Name()] = z.String
			}
		}

		column["model_name"] = importer.columnNameMap(column["sql_name"])
		if v, ok := importer.schemaMap[column["sql_type"]]; ok {
			column["model_type"] = v
		} else {
			return nil, fmt.Errorf("sql type %s not in schema map", column["sql_type"])
		}

		columns = append(columns, column)
	}

	tColumns := []importedColumn{}

	byt, err := json.Marshal(&columns)

	if err != nil {
		return nil, fmt.Errorf("marshalling column: %w", err)
	}

	err = json.Unmarshal(byt, &tColumns)

	if err != nil {
		return nil, fmt.Errorf("unmarshalling column: %w", err)
	}

	return tColumns, nil
}

func NewImporter(schemaName string, outputDest string, ignoredTables []string) *_importer {
	return &_importer{
		outputDest:    outputDest,
		schemaName:    schemaName,
		columnNameMap: DefaultNameMap,
		schemaMap:     DefaultSchemaMap(),
		ignoredTables: ignoredTables,
	}
}
