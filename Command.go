package gormy

import (
	"fmt"

	"github.com/radasam/gormy/internal/driver"
)

type Command struct {
	queryString string
	errored     error
}

func (command *Command) Exec() (driver.CommandResult, error) {

	if command.errored != nil {
		return nil, command.errored
	}

	queryString := command.queryString

	result, err := gc.conn.Exec(queryString)

	if err != nil {
		return nil, fmt.Errorf("executing command: %w", err)
	}

	return result, nil
}
