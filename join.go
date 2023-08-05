package gormy

import (
	"github.com/radasam/gormy/internal/driver"
)

type Join interface {
	Columns() []Column
	ColumnExpr() string
	TableExpr() string
	JoinExpr(originKey string, relation Relation) string
	Parse(parentRow string, key string, name string, column driver.ColumnType, sqlType interface{})
	Row(parentRow string) interface{}
	Values() interface{}
	JoinName() string
	JoinKey() string
	OnJoin(join Join)
	HasJoin() bool
}
