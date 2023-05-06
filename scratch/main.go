package main

import (
	"encoding/json"
	"gormy/test/models"
	"reflect"
)

func main() {
	println("hello")

	mytable := models.MyTable{}

	fields := reflect.TypeOf(mytable)

	println(fields.Field(3).Type.Field(0).Tag)

	mytable2 := models.MySecondTable{Age: 22, Name: "sam", Color: "green"}

	bytes, _ := json.Marshal(mytable2)

	mytable3 := reflect.New(fields.Field(3).Type)

	c := mytable3.Elem().Interface().(models.MySecondTable)

	_ = json.Unmarshal(bytes, &c)

	println(c.Name)
}
