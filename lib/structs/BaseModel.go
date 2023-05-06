package structs

import (
	"gormy/lib/modelparser"
	"gormy/lib/types"
	"reflect"
)

type BaseModel interface {
}

func Model[T any](baseModel T) *Table[T] {

	myTableDDL := &Table[T]{}

	fields := reflect.TypeOf(baseModel)

	if fields.Kind() == reflect.Slice {
		fields = fields.Elem()
	}

	myTableDDL.Columns = []types.Column{}
	myTableDDL.Relations = []types.Relation{}
	myTableDDL.Rows = []T{baseModel}
	relationCount := 0

	for i := 0; i < fields.NumField(); i++ {
		if fields.Field(i).Name != "baseModel" {
			column, relation, err := modelparser.ParseColumn(string(fields.Field(i).Tag), fields.Field(i).Name, fields.Field(i).Type, "jk0", len(myTableDDL.Relations))

			if err != nil {
				println(err.Error())
			}

			if column != nil {
				myTableDDL.Columns = append(myTableDDL.Columns, *column)
			}

			if relation != nil {
				myTableDDL.Relations = append(myTableDDL.Relations, *relation)
				relationCount += 1
			}

		} else {
			tableName, err := modelparser.ParseConfig(string(fields.Field(i).Tag))

			if err != nil {
				println(err.Error())
			}

			myTableDDL.Name = tableName
		}
	}

	return myTableDDL
}
