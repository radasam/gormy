package models

import (
	"github.com/radasam/gormy/pkg/gormy"
)

type MySecondTable struct {
	baseModel gormy.BaseModel `gormy:"mysecondtable"`
	Name      string          `gormy:"varchar"`
	Age       int             `gormy:"int,name:age"`
	Color     string          `gormy:"varchar"`
}
