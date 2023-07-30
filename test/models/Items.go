package models

import (
	"github.com/radasam/gormy/pkg/gormy"
)

type Items struct {
	baseModel gormy.BaseModel `gormy:"items"`
	ItemId    string          `gormy:"varchar,name:item_id"`
	OrderId   string          `gormy:"varchar,name:order_id"`
	Name      string          `gormy:"varchar"`
}
