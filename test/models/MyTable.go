package models

type MyTable struct {
	Name string `gormy:"varchar"`
	Age  int    `gormy:"int,name:age"`
}
