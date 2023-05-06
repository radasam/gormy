package models

import "gormy/lib/structs"

type Orders struct {
	baseModel structs.BaseModel `gormy:"orders"`
	ItemId    string            `gormy:"varchar,name:item_id"`
	Timestamp int               `gormy:"int"`
}
