package models

import (
	"github.com/radasam/gormy/lib/engine"
)

type MySecondTable struct {
	baseModel engine.BaseModel `gormy:"mysecondtable"`
	Name      string           `gormy:"varchar"`
	Age       int              `gormy:"int,name:age"`
	Color     string           `gormy:"varchar"`
}
