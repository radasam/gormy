package gormy

import (
	"database/sql"
)

type Join interface {
	Columns() []Column
	ColumnExpr() string
	TableExpr() string
	JoinExpr(originKey string, relation Relation) string
	Parse(parentRow string, key string, name string, column *sql.ColumnType, sqlType interface{})
	Row(parentRow string) interface{}
	Values() interface{}
	JoinName() string
	JoinKey() string
	OnJoin(join Join)
}
