package models

import "gormy/lib/structs"

type Items struct {
	baseModel structs.BaseModel `gormy:"items"`
	Id        string            `gormy:"varchar"`
	Name      string            `gormy:"varchar"`
	Orders    []Orders          `gormy:"relation:onetomany,how:left,on:id=item_id"`
}
