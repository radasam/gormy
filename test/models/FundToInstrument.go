package models

import (
	"github.com/radasam/gormy/pkg/gormy"
)

type FundToInstrument struct {
	baseModel    gormy.BaseModel `gormy:"fund_to_instrument"`
	FundId       string          `gormy:"varchar,name:fund_id"`
	InstrumentId string          `gormy:"varchar,name:instrument_id"`
}
