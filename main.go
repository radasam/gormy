package main

import (
	"gormy/lib/joins"
	"gormy/lib/structs"
	"gormy/test/models"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	joins.Init()
	// myTable := models.MyTable{
	// 	Name: "Sam",
	// 	Age:  26,
	// }

	// myTableDDL := structs.Table[models.MyTable]{}

	// fields := reflect.TypeOf(myTable)

	// myTableDDL.Name = "MyTable"
	// myTableDDL.Columns = []structs.Column{}

	// for i := 0; i < fields.NumField(); i++ {
	// 	println(fields.Field(i).Tag)
	// 	column, err := modelparser.ParseColumn(string(fields.Field(i).Tag), fields.Field(i).Name)

	// 	if err != nil {
	// 		println(err.Error())
	// 	}

	// 	myTableDDL.Columns = append(myTableDDL.Columns, column)
	// }

	// _, err := engine.DB().Exec(myTableDDL.Create())

	// if err != nil {
	// 	println(err.Error())
	// }

	// result := myTableDDL.Query().Select().Column("age").Column("name").Where("? = ?", "age", "26").Exec()

	// for _, mt := range result {
	// 	println(mt.Age)
	// 	println(mt.Name)
	// }

	// result := structs.Model(models.MyTable{}).Query().Select().Column("age").Column("name").Where("? = ?", "age", "26").Exec()

	// structs.Model(models.MyTable{}).Create()

	// for _, mt := range result {
	// 	println(mt.Age)
	// 	println(mt.Name)
	// }

	// myTable := structs.Model(models.MyTable{}).Query().Select().Relation("SecondTable", "onetoone").Exec()

	// spew.Dump(myTable)

	// items := structs.Model(models.Items{}).Query().Select().Relation("Orders", "onetomany").Exec()

	// for _, item := range items {
	// 	println(item.Id)
	// 	for _, o := range item.Orders {
	// 		println(fmt.Sprintf("%s  %s abc", item.Id, o.ItemId))
	// 	}
	// }

	// spew.Dump(items)

	// funds := []models.Funds{{Name: "fund-1", FundId: "F01", Instruments: []models.Instruments{{InstrumentId: "I02", Name: "instrument-2"}, {InstrumentId: "I04", Name: "instrument-4"}}}, {FundId: "F02", Name: "fund-2", Instruments: []models.Instruments{{InstrumentId: "I01", Name: "instrument-1"}, {InstrumentId: "I02", Name: "instrument-2"}, {InstrumentId: "I03", Name: "instrument-3"}}}}

	// _, err := structs.Model(models.Funds{}).Query().Insert(&funds).Exec()

	// if err != nil {
	// 	println(err.Error())
	// }

	// fundsToInstruments := []models.FundToInstrument{{FundId: "F01", InstrumentId: "I02"}, {FundId: "F01", InstrumentId: "I04"}, {FundId: "F02", InstrumentId: "I01"}, {FundId: "F02", InstrumentId: "I02"}, {FundId: "F02", InstrumentId: "I03"}}

	// _, err = structs.Model(models.FundToInstrument{}).Query().Insert(&fundsToInstruments).Exec()

	// if err != nil {
	// 	println(err.Error())
	// }

	// Instruments := []models.Instruments{{InstrumentId: "I01", Name: "Instrument-1"}, {InstrumentId: "I02", Name: "Instrument-2"}, {InstrumentId: "I03", Name: "Instrument-3"}, {InstrumentId: "I04", Name: "Instrument-4"}}

	// _, err = structs.Model(models.Instruments{}).Query().Insert(&Instruments).Exec()

	// if err != nil {
	// 	println(err.Error())
	// }

	funds := structs.Model(models.Funds{}).Query().Select().Relation("Instruments", "manytomany").Exec()

	spew.Dump(funds[0].Instruments[0])
}
