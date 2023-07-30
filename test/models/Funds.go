package models

import (
	"github.com/radasam/gormy/pkg/gormy"
)

type Funds struct {
	baseModel   gormy.BaseModel `gormy:"funds"`
	FundId      string          `gormy:"varchar,name:fund_id"`
	Name        string          `gormy:"varchar,name:name"`
	Instruments []Instruments   `gormy:"relation:manytomany,how:left,on:fund_id=instrument_id,intermediary:fund_to_instrument"`
}
