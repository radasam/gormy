package models

import (
	"github.com/radasam/gormy/pkg/gormy"
)

type Orders struct {
	baseModel gormy.BaseModel `gormy:"orders"`
	OrderId   string          `gormy:"varchar,name:order_id"`
	UserId    string          `gormy:"varchar,name:user_id"`
	Timestamp int             `gormy:"int"`
	Items     []Items         `gormy:"relation:onetomany,how:left,on:order_id=order_id"`
}
