package gormy

import "database/sql"

type JoinParser interface {
	Parse(parentRow string, key string, name string, column *sql.ColumnType, sqlType interface{})
	Row(parentRow string) interface{}
	Values() interface{}
	OnJoin(join Join)
}
