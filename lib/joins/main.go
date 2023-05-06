package joins

import (
	"database/sql"
	"fmt"
	"gormy/lib/types"
)

type Join interface {
	Columns() []types.Column
	ColumnExpr() string
	TableExpr() string
	JoinExpr(relation types.Relation) string
	Parser(rowNumber int, repeatRowNumber int, key string, name string, column *sql.ColumnType, sqlType interface{})
	Values(rowNumber int) interface{}
	JoinName() string
	JoinKey() string
	OnJoin()
}

type joins struct {
	JoinMap map[string]func(joinKey string, joinName string, joinsTo string, columns []types.Column, tableExpr string) Join
}

func (joins *joins) Register(name string, join func(joinKey string, joinName string, joinsTo string, columns []types.Column, tableExpr string) Join) {
	joins.JoinMap[name] = join
}

func (joins *joins) ByName(name string) (func(joinKey string, joinName string, joinsTo string, columns []types.Column, tableExpr string) Join, error) {
	for k := range joins.JoinMap {
		if k == name {
			return joins.JoinMap[k], nil
		}
	}

	return nil, fmt.Errorf("Join doesnt exist")

}

var Joins *joins

func Init() {
	Joins = &joins{
		JoinMap: map[string]func(joinkey string, joinName string, joinsTo string, columns []types.Column, tableExpr string) Join{},
	}
	Joins.Register("onetoone", OneToOne)
	Joins.Register("onetomany", OneToMany)
	Joins.Register("manytomany", ManyToMany)
}
