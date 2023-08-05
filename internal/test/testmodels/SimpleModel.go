package testmodels

import "github.com/radasam/gormy"

type SimpleModel struct {
	baseModel gormy.BaseModel `gormy:"simplemodel"`
	Name      string          `gormy:"varchar"`
	Age       int             `gormy:"int"`
}
