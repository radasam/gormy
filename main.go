package gormy

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	// "github.com/radasam/gormy/internal/joins"
)

var gc *GormyClient

func db() *sql.DB {
	if gc == nil {
		panic("GormyClient has not been initialised!")
	}
	return gc.conn
}

type GormyClient struct {
	conn *sql.DB
}

// type Join joins.Join

// func (gc *GormyClient) RegisterJoin(name string, join func(joinKey string, joinName string, joinsTo string, columns []Column, tableExpr string, parentJoinRow string) joins.Join) error {
// 	joins.Joins.Register(name, join)
// }

func NewGormyClient(connString string) (*GormyClient, error) {
	conn, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, fmt.Errorf("connecting to db: %w", err)
	}

	gc = &GormyClient{
		conn: conn,
	}

	return gc, nil
}
