package models

import (
	"github.com/radasam/gormy/lib/engine"
)

type MyTable struct {
	baseModel   engine.BaseModel `gormy:"mytable"`
	Name        string           `gormy:"varchar"`
	Age         int              `gormy:"int,name:age"`
	SecondTable MySecondTable    `gormy:"relation:onetoone,how:left,on:name=name"`
}
