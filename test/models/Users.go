package models

import (
	"github.com/radasam/gormy/pkg/gormy"
)

type Users struct {
	baseModel gormy.BaseModel `gormy:"users"`
	UserId    string          `gormy:"varchar,name:user_id"`
	UserName  string          `gormy:"varchar"`
	Orders    []Orders        `gormy:"relation:onetomany,how:left,on:user_id=user_id"`
}
