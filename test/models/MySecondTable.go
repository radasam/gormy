package models

import "gormy/lib/structs"

type MySecondTable struct {
	baseModel structs.BaseModel `gormy:"mysecondtable"`
	Name      string            `gormy:"varchar"`
	Age       int               `gormy:"int,name:age"`
	Color     string            `gormy:"varchar"`
}
