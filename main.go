package gormy

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/radasam/gormy/internal/driver"
)

type registeredjoins struct {
	JoinMap map[string]func(joinKey string, joinName string, joinsTo string, columns []Column, tableExpr string, parentJoinRow string) Join
}

func (rj *registeredjoins) Register(name string, join func(joinKey string, joinName string, joinsTo string, columns []Column, tableExpr string, parentJoinRow string) Join) {
	rj.JoinMap[name] = join
}

var gc *GormyClient

func db() driver.Driver {
	if gc == nil {
		panic("GormyClient has not been initialised!")
	}
	return gc.conn
}

type GormyClient struct {
	conn            driver.Driver
	RegisteredJoins *registeredjoins
}

// type Join joins.Join

// func (gc *GormyClient) RegisterJoin(name string, join func(joinKey string, joinName string, joinsTo string, columns []Column, tableExpr string, parentJoinRow string) joins.Join) error {
// 	joins.Joins.Register(name, join)
// }

func NewGormyClient(connString string) (*GormyClient, error) {
	conn, err := driver.NewPostgres(connString)

	if err != nil {
		return nil, fmt.Errorf("connecting to db: %w", err)
	}

	registeredJoins := &registeredjoins{
		JoinMap: map[string]func(joinkey string, joinName string, joinsTo string, columns []Column, tableExpr string, parentJoinRow string) Join{},
	}
	registeredJoins.Register("onetoone", OneToOne)
	registeredJoins.Register("onetomany", OneToMany)
	registeredJoins.Register("manytomany", ManyToMany)

	gc = &GormyClient{
		conn:            conn,
		RegisteredJoins: registeredJoins,
	}

	return gc, nil
}

func UseMockClient(conn *driver.Mock) (*GormyClient, error) {

	registeredJoins := &registeredjoins{
		JoinMap: map[string]func(joinkey string, joinName string, joinsTo string, columns []Column, tableExpr string, parentJoinRow string) Join{},
	}
	registeredJoins.Register("onetoone", OneToOne)
	registeredJoins.Register("onetomany", OneToMany)
	registeredJoins.Register("manytomany", ManyToMany)

	gc = &GormyClient{
		conn:            conn,
		RegisteredJoins: registeredJoins,
	}

	return gc, nil
}

func (rj *registeredjoins) ByName(name string) (func(joinKey string, joinName string, joinsTo string, columns []Column, tableExpr string, parentJoinRow string) Join, error) {
	for k := range rj.JoinMap {
		if k == name {
			return rj.JoinMap[k], nil
		}
	}

	return nil, fmt.Errorf("Join doesnt exist")

}

func Init() {

}

func RegisterJoin(name string, join func(joinKey string, joinName string, joinsTo string, columns []Column, tableExpr string, parentJoinRow string) Join) {
	gc.RegisteredJoins.Register(name, join)
}
