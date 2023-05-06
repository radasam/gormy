package sqlparser

import (
	"database/sql"
	"encoding/json"
	"gormy/lib/joins"
	"strings"
)

type sqlParser[T any] struct {
	rows   sql.Rows
	joins  []joins.Join
	values map[int]map[string]interface{}
}

func (sqlparser *sqlParser[T]) Parser(rowNumber int, name string, column *sql.ColumnType, sqlType interface{}) {

	if z, ok := (sqlType).(*sql.NullBool); ok {
		sqlparser.values[rowNumber][name] = z.Bool
		return
	}

	if z, ok := (sqlType).(*sql.NullString); ok {
		sqlparser.values[rowNumber][name] = z.String
		return
	}

	if z, ok := (sqlType).(*sql.NullInt64); ok {
		sqlparser.values[rowNumber][name] = z.Int64
		return
	}

	if z, ok := (sqlType).(*sql.NullFloat64); ok {
		sqlparser.values[rowNumber][name] = z.Float64
		return
	}

	if z, ok := (sqlType).(*sql.NullInt32); ok {
		sqlparser.values[rowNumber][name] = z.Int32
		return
	}

	sqlparser.values[rowNumber][name] = sqlType
}

func (sqlparser *sqlParser[T]) toJson() (string, error) {
	sqlparser.values = map[int]map[string]interface{}{}
	columnTypes, err := sqlparser.rows.ColumnTypes()

	if err != nil {
		return "", err
	}

	count := len(columnTypes)
	lastRowNumber := 0
	rowNumber := 0
	repeatRowNumber := 0

	for sqlparser.rows.Next() {

		scanArgs := make([]interface{}, count)

		for i, v := range columnTypes {
			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
				break
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
				break
			case "INT4", "INT8":
				scanArgs[i] = new(sql.NullInt64)
				break
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		err := sqlparser.rows.Scan(scanArgs...)

		if err != nil {
			return "", err
		}

		if columnTypes[0].Name() == "jk0__join_row" {
			if z, ok := (scanArgs[0]).(*sql.NullInt64); ok {
				rowNumber = int(z.Int64)
			}
		} else {
			rowNumber += 1
		}

		if lastRowNumber == rowNumber {
			repeatRowNumber += 1
		} else {
			lastRowNumber = rowNumber
			repeatRowNumber = 0
		}

		if _, ok := sqlparser.values[rowNumber]; !ok {
			sqlparser.values[rowNumber] = map[string]interface{}{}
		}

		for i, column := range columnTypes {
			if column.Name() != "jk0__join_row" {
				key := strings.Split(column.Name(), "__")[0]
				name := strings.Split(column.Name(), "__")[1]

				if key == "jk0" {
					sqlparser.Parser(rowNumber, name, column, scanArgs[i])
				} else {
					for _, join := range sqlparser.joins {
						join.Parser(rowNumber, repeatRowNumber, key, name, column, scanArgs[i])
					}
				}
			}

		}
		for _, v := range sqlparser.joins {
			sqlparser.values[rowNumber][v.JoinName()] = v.Values(rowNumber)
		}
	}

	values := make([]map[string]interface{}, 0, len(sqlparser.values))
	for k := range sqlparser.values {
		values = append(values, sqlparser.values[k])
	}

	z, err := json.Marshal(values)

	if err != nil {
		return "", err
	}

	return string(z), err
}

func (sqlparser *sqlParser[T]) Parse(rows *T) error {
	jsonString, err := sqlparser.toJson()

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(jsonString), &rows)

	if err != nil {
		return err
	}

	return nil

}

func NewSqlParser[T any](rowStruct T, joins []joins.Join, rows sql.Rows) sqlParser[T] {
	return sqlParser[T]{
		joins: joins,
		rows:  rows,
	}
}
