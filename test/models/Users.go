package models

import "gormy/lib/structs"

type Users struct {
	baseModel structs.BaseModel `gormy:"users"`
	UserId    string            `gormy:"varchar,name:user_id"`
	UserName  string            `gormy:"varchar"`
	Orders    []Orders          `gormy:"relation:onetomany,how:left,on:user_id=user_id"`
}
