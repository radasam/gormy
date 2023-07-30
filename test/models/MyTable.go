package models

import (
	"github.com/radasam/gormy/pkg/gormy"
)

type MyTable struct {
	baseModel   gormy.BaseModel `gormy:"mytable"`
	Name        string          `gormy:"varchar"`
	Age         int             `gormy:"int,name:age"`
	SecondTable MySecondTable   `gormy:"relation:onetoone,how:left,on:name=name"`
}
