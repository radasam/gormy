package models

import (
	"github.com/radasam/gormy/lib/engine"
)

type Instruments struct {
	baseModel    engine.BaseModel `gormy:"instruments"`
	InstrumentId string           `gormy:"varchar,name:instrument_id"`
	Name         string           `gormy:"varchar,name:name"`
	Funds        []Funds          `gormy:"relation:manytomany,how:left,on:instrument_id=instrument_id,intermediary:fund_to_instrument"`
}
