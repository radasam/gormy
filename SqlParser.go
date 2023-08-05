package gormy

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/radasam/gormy/internal/driver"
)

type sqlParser[T any] struct {
	rows   driver.RowsResult
	values map[int]map[string]interface{}
	origin Join
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
		return "", fmt.Errorf("failed to read column types from sql result: %w", err)
	}

	count := len(columnTypes)
	rowNumber := 0

	for sqlparser.rows.Next() {

		scanArgs := make([]interface{}, count)

		for i, v := range columnTypes {
			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
			case "INT4", "INT8":
				scanArgs[i] = new(sql.NullInt64)
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		err := sqlparser.rows.Scan(scanArgs...)

		if err != nil {
			return "", fmt.Errorf("failed to scan sql result: %w", err)
		}

		if columnTypes[0].Name() == "jk0__join_row" {
			if z, ok := (scanArgs[0]).(*sql.NullInt64); ok {
				rowNumber = int(z.Int64) - 1
			}
		}

		for i, column := range columnTypes {

			key := strings.Split(column.Name(), "__")[0]
			name := strings.Split(column.Name(), "__")[1]
			sqlparser.origin.Parse(strconv.Itoa(rowNumber), key, name, column, scanArgs[i])
		}
	}

	values := sqlparser.origin.Values()

	z, err := json.Marshal(&values)

	if err != nil {
		return "", fmt.Errorf("marshalling result to struct: %w", err)
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
		return fmt.Errorf("failed to unmarshal result to struct: %w", err)
	}

	return nil

}

func newSqlParser[T any](rowStruct T, origin Join, rows driver.RowsResult) sqlParser[T] {
	return sqlParser[T]{
		origin: origin,
		rows:   rows,
	}
}
