package gormy

import (
	"fmt"
	"reflect"
)

type BaseModel interface {
}

func Model[T any](baseModel T) *Table[T] {

	table := &Table[T]{}

	fields := reflect.TypeOf(baseModel)

	if fields.Kind() == reflect.Slice {
		fields = fields.Elem()
	}

	table.Columns = []Column{}
	table.Relations = []Relation{}
	table.Rows = []T{baseModel}
	relationCount := 0

	for i := 0; i < fields.NumField(); i++ {
		if fields.Field(i).Name != "baseModel" {
			column, relation, err := ParseColumn(string(fields.Field(i).Tag), fields.Field(i).Name, fields.Field(i).Type, "jk0", len(table.Relations))

			if err != nil {
				table.errored = fmt.Errorf("parsing column: %w", err)
			}

			if column != nil {
				table.Columns = append(table.Columns, *column)
			}

			if relation != nil {
				table.Relations = append(table.Relations, *relation)
				relationCount += 1
			}

		} else {
			tableName, err := ParseConfig(string(fields.Field(i).Tag))

			if err != nil {
				table.errored = fmt.Errorf("parsing config: %w", err)
			}

			table.Name = tableName
		}
	}

	return table
}
