package gormy

import "github.com/radasam/gormy/internal/driver"

type JoinParser interface {
	Parse(parentRow string, key string, name string, column driver.ColumnType, sqlType interface{})
	Row(parentRow string) interface{}
	Values() interface{}
	OnJoin(join Join)
}
