package models

import (
	"github.com/radasam/gormy/lib/engine"
)

type Users struct {
	baseModel engine.BaseModel `gormy:"users"`
	UserId    string           `gormy:"varchar,name:user_id"`
	UserName  string           `gormy:"varchar"`
	Orders    []Orders         `gormy:"relation:onetomany,how:left,on:user_id=user_id"`
}
