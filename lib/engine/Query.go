package engine

import (
	"fmt"
	"gormy/lib/joins"
	"gormy/lib/types"
	"gormy/lib/utils"
	"reflect"
)

type Query[T any] struct {
	origin      joins.Join
	tableName   string
	columns     []types.Column
	queryString string
	Rows        []T
	relations   []types.Relation
}

func (query *Query[T]) Select() *SelectQuery[T] {
	newQuery := &SelectQuery[T]{}

	newQuery.origin = query.origin
	newQuery.tableName = query.tableName
	newQuery.queryString = "SELECT $columns$ \r\nFROM $jk0__table_name$ as jk0\r\n"

	newQuery.columns = query.columns
	newQuery.relations = query.relations
	newQuery.activeRelations = make([]ActiveRelation, 0)

	return newQuery
}

func (query *Query[T]) Insert(rows *[]T) *Command {

	if rows != nil {
		query.Rows = *rows
	}

	command := &Command{}

	command.queryString = fmt.Sprintf("INSERT INTO %s (", query.tableName)

	validColumns := []int{}

	for columnIndex, column := range query.columns {
		if !column.IsRelation {
			if columnIndex != len(query.columns)-1 {
				command.queryString += fmt.Sprintf("%s ,", column.Name)
			} else {
				command.queryString += fmt.Sprintf("%s)\r\n", column.Name)
			}
			validColumns = append(validColumns, columnIndex+1)
		}
	}

	command.queryString += "VALUES\r\n"

	for rowIndex, row := range query.Rows {
		fields := reflect.ValueOf(row)

		for i, v := range validColumns {
			if i == 0 {
				command.queryString += "	(" + utils.StructToRow(fields.Field(v)) + ", "
			} else if i == len(validColumns)-1 {
				command.queryString += utils.StructToRow(fields.Field(v)) + ")"
			} else {
				command.queryString += utils.StructToRow(fields.Field(v)) + ", "
			}
		}

		if rowIndex != len(query.Rows)-1 {
			command.queryString += ",\r\n"
		}
	}

	return command

}

func (query *Query[T]) columnByName(columnName string) (types.Column, error) {
	for _, column := range query.origin.Columns() {
		if column.Name == columnName {
			return column, nil
		}
	}

	return types.Column{}, fmt.Errorf("Column doesnt exist")
}
