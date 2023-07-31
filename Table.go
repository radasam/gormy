package gormy

import (
	"fmt"
)

type Table[T any] struct {
	Name      string
	Columns   []Column
	Relations []Relation
	Rows      []T
}

func (table *Table[T]) Create() *Command {
	createStatement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ( \r\n", table.Name)

	for columnIndex, column := range table.Columns {
		if columnIndex != len(table.Columns)-1 {
			createStatement += fmt.Sprintf("%s %s,\r\n", column.Name, column.DataType)
		} else {
			createStatement += fmt.Sprintf("%s %s\r\n", column.Name, column.DataType)
		}
	}

	createStatement += ");"

	return &Command{
		queryString: createStatement,
	}
}

func (table *Table[T]) Query() *Query[T] {
	newQuery := &Query[T]{}

	newQuery.tableName = table.Name
	newQuery.columns = table.Columns
	newQuery.relations = table.Relations
	newQuery.Rows = table.Rows
	newQuery.origin = Origin("Origin", table.Columns, table.Name)

	return newQuery
}
