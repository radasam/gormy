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
	JoinExpr(originKey string, relation types.Relation) string
	Parse(parentRow string, key string, name string, column *sql.ColumnType, sqlType interface{})
	Row(parentRow string) interface{}
	Values() interface{}
	JoinName() string
	JoinKey() string
	OnJoin(join Join)
}

type joins struct {
	JoinMap map[string]func(joinKey string, joinName string, joinsTo string, columns []types.Column, tableExpr string, parentJoinRow string) Join
}

func (joins *joins) Register(name string, join func(joinKey string, joinName string, joinsTo string, columns []types.Column, tableExpr string, parentJoinRow string) Join) {
	joins.JoinMap[name] = join
}

func (joins *joins) ByName(name string) (func(joinKey string, joinName string, joinsTo string, columns []types.Column, tableExpr string, parentJoinRow string) Join, error) {
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
		JoinMap: map[string]func(joinkey string, joinName string, joinsTo string, columns []types.Column, tableExpr string, parentJoinRow string) Join{},
	}
	Joins.Register("onetoone", OneToOne)
	Joins.Register("onetomany", OneToMany)
	Joins.Register("manytomany", ManyToMany)
}
