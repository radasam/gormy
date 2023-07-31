package gormy

import (
	"database/sql"
	"fmt"
)

type Command struct {
	queryString string
	errored     error
}

func (command *Command) Exec() (sql.Result, error) {

	if command.errored != nil {
		return nil, command.errored
	}

	queryString := command.queryString

	result, err := gc.conn.Exec(queryString)

	return result, fmt.Errorf("executing command: %w", err)
}
