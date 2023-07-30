package main

import (
	"fmt"
	"gormy/lib/engine"
	"gormy/lib/joins"
	"gormy/test/models"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	joins.Init()
	_, err := engine.NewGormyClient(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres", "postgrespw",
		"localhost", "55000", "postgres"))

	if err != nil {
		println(err.Error())
		return
	}
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

	// funds := structs.Model(models.Funds{}).Query().Select().Relation("Instruments", "manytomany").Exec()

	// spew.Dump(funds[0].Instruments[0])

	// structs.Model(models.Users{}).Create().Exec()
	// structs.Model(models.Orders{}).Create().Exec()
	// structs.Model(models.Items{}).Create().Exec()

	// items := []models.Items{
	// 	{ItemId: "item-123", OrderId: "order-456", Name: "spoon"},
	// 	{ItemId: "item-234", OrderId: "order-456", Name: "fork"},
	// 	{ItemId: "item-345", OrderId: "order-789", Name: "knife"},
	// 	{ItemId: "item-456", OrderId: "order-789", Name: "whisk"},
	// 	{ItemId: "item-567", OrderId: "order-123", Name: "spatula"},
	// 	{ItemId: "item-678", OrderId: "order-123", Name: "ladle"},
	// 	{ItemId: "item-789", OrderId: "order-123", Name: "tongs"},
	// 	{ItemId: "item-890", OrderId: "order-234", Name: "grater"},
	// 	{ItemId: "item-901", OrderId: "order-234", Name: "peeler"},
	// 	{ItemId: "item-012", OrderId: "order-234", Name: "can opener"},
	// }

	// structs.Model(models.Items{}).Query().Insert(&items).Exec()

	// orders := []models.Orders{
	// 	{OrderId: "order-123", UserId: "user-001", Timestamp: 1620418507},
	// 	{OrderId: "order-234", UserId: "user-001", Timestamp: 1620418523},
	// 	{OrderId: "order-456", UserId: "user-002", Timestamp: 1620418534},
	// 	{OrderId: "order-789", UserId: "user-002", Timestamp: 1620418551},
	// }

	// structs.Model(models.Orders{}).Query().Insert(&orders).Exec()

	// users := []models.Users{
	// 	{UserId: "user-001", UserName: "Gordon Ramsay"},
	// 	{UserId: "user-002", UserName: "Julia Child"},
	// }

	// structs.Model(models.Users{}).Query().Insert(&users).Exec()

	myUsers := engine.Model(models.Users{}).Query().Select().Relation("Orders", "onetomany").Relation("Items", "onetomany").Exec()

	spew.Dump(myUsers[0].Orders[0])

	// myImporter := engine.NewImporter("public", "./test-import")

	// err = myImporter.Import()

	// if err != nil {
	// 	println(err.Error())
	// }

}
