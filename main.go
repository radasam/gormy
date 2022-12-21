package main

import (
	"gormy/lib/modelparser"
	"gormy/lib/structs"
	"gormy/test/models"
	"reflect"
)

func main() {
	myTableDDL := structs.Table{}

	myTable := models.MyTable{
		Name: "Sam",
		Age:  26,
	}

	fields := reflect.TypeOf(myTable)

	myTableDDL.Name = "MyTable"

	for i := 0; i < fields.NumField(); i++ {
		println(fields.Field(i).Tag)
		column, err := modelparser.ParseColumn(string(fields.Field(i).Tag), fields.Field(i).Name)

		if err != nil {
			println(err.Error())
		}

		println(column.Name)
		println(column.DataType)
	}

}
