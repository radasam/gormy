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
type MyTable struct {
	baseModel   gormy.BaseModel `gormy:"mytable"`
	Name        string          `gormy:"varchar"`
	Age         int             `gormy:"int,name:age"`
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

```postgres://user:pass@host:port/schema``

