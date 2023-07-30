package gormy

import (
	"database/sql"
)

type Command struct {
	queryString string
}

func (command *Command) Exec() (sql.Result, error) {
	queryString := command.queryString

	result, err := db().Exec(queryString)

	return result, err
}
