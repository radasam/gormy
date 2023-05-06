package models

import "gormy/lib/structs"

type MyTable struct {
	baseModel   structs.BaseModel `gormy:"mytable"`
	Name        string            `gormy:"varchar"`
	Age         int               `gormy:"int,name:age"`
	SecondTable MySecondTable     `gormy:"relation:onetoone,how:left,on:name=name"`
}
