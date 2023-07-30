package models

import (
	"github.com/radasam/gormy/lib/engine"
)

type Items struct {
	baseModel engine.BaseModel `gormy:"items"`
	ItemId    string           `gormy:"varchar,name:item_id"`
	OrderId   string           `gormy:"varchar,name:order_id"`
	Name      string           `gormy:"varchar"`
}
