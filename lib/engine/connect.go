package engine

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DB() *sql.DB {
	if db == nil {
		connStrnig := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			"postgres", "postgrespw",
			"localhost", "55000", "postgres")
		// fmt.Println(connStrnig)
		conn, err := sql.Open("postgres", connStrnig)
		if err != nil {
			panic(err)
		}
		return conn
	}
	return db
}
