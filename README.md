# Gormy

A postgres ORM for go.

## Installation

```bash
go get github.com/radasam/gormy
```

## Examples

### Defining a model

To make a query you first need to define a database model

```go
type MyFirstModel struct {
	baseModel gormy.BaseModel `gormy:"myfirstmodel"`
	Name      string          `gormy:"varchar"`
	Age       int             `gormy:"int,name:age"`
	Color     string          `gormy:"varchar"`
}

```

A database model is just a go struct with some struct tags, each model must have the baseModel attribute with a struct tag that corresponds to the table name in the database.

To define table columns add attributes to the struct with a tag consisting of the datatype of the column and optionally the name it appears with in the database.

### Initialising a client

Before making a query we must also initialise the gormy client

```go
gc, err := gormy.NewGormyClient(myConnString)
```

Where myConnString is a postgres connection string of the form

```postgres://user:pass@host:port/schema```

### Performing a query

Now we have the model and client we can make a query

```
firstModel, err := gormy.Model(models.MyFirstModel{}).Query().Select().Where("? = '?'", "name", "Steve").Exec()
```

### Creating a table

You can also use models to define and create new tables in your data base.

```go
_, err = gormy.Model(models.MyFirstModel{}).Create().Exec()
```

### Inserting data

If you pass a struct with values to Model you can also insert data

```go
myfirstmodel := []models.MyFirstModel{
	{
		Name:  "Steve",
		Age:   50,
		Color: "blue",
	},
	{
		Name:  "Mary",
		Age:   27,
		Color: "green",
	},
}

_, err = gormy.Model(models.MyFirstModel{}).Query().Insert(&myfirstmodel).Exec()

```

### Joins

To use a join on a table you first must add the join to the model, specifying a join type, and a column to join on.

```go
type MyFirstModel struct {
	baseModel   gormy.BaseModel `gormy:"myfirstmodel"`
	Name        string          `gormy:"varchar"`
	Age         int             `gormy:"int,name:age"`
	Color       string          `gormy:"varchar"`
	SecondTable MySecondTable   `gormy:"relation:onetoone,how:left,on:name=name"`
}
```
You can then use the join as part of a query

```go
firstModel, err := gormy.Model(models.MyTable{}).Query().Select().Relation("SecondTable", "onetoone").Where("? = '?'", "Name", "sam").Exec()
```

### Importing tables

If you already have your tables defined in your database you can import the database tables into gormy models.

```go
err = gormy.NewImporter("myschema", "./myoutputdir", []string{"ignorethistable"}).Import()
```