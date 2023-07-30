package models

import (
	"github.com/radasam/gormy/lib/engine"
)

type FundToInstrument struct {
	baseModel    engine.BaseModel `gormy:"fund_to_instrument"`
	FundId       string           `gormy:"varchar,name:fund_id"`
	InstrumentId string           `gormy:"varchar,name:instrument_id"`
}
