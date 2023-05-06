package models

import "gormy/lib/structs"

type FundToInstrument struct {
	baseModel    structs.BaseModel `gormy:"fund_to_instrument"`
	FundId       string            `gormy:"varchar,name:fund_id"`
	InstrumentId string            `gormy:"varchar,name:instrument_id"`
}
