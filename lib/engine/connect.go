package engine

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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
