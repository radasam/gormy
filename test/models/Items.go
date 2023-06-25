package models

import "gormy/lib/structs"

type Items struct {
	baseModel structs.BaseModel `gormy:"items"`
	ItemId    string            `gormy:"varchar,name:item_id"`
	OrderId   string            `gormy:"varchar,name:order_id"`
	Name      string            `gormy:"varchar"`
}
