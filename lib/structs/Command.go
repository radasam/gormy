package structs

import (
	"database/sql"
	"gormy/lib/engine"
)

type Command struct {
	queryString string
}

func (command *Command) Exec() (sql.Result, error) {
	queryString := command.queryString

	result, err := engine.DB().Exec(queryString)

	return result, err
}
