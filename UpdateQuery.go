package gormy

import (
	"fmt"
	"strings"

	"github.com/radasam/gormy/internal/driver"
)

type UpdateQuery[T any] struct {
	Query[T]
	errored error
	set     []Statement
	where   []Statement
}

func (uq *UpdateQuery[T]) Set(expr string, columnName string, value string) *UpdateQuery[T] {

	uq.set = append(uq.set, Statement{expr, columnName, value})

	return uq
}

func (uq *UpdateQuery[T]) Where(expr string, columnName string, value string) *UpdateQuery[T] {
	uq.where = append(uq.where, Statement{expr, columnName, value})

	return uq
}

func (uq *UpdateQuery[T]) Exec() (driver.CommandResult, error) {

	for i, s := range uq.set {
		if i == 0 {
			uq.queryString += fmt.Sprintf("SET %s", s.ToString())
		} else {
			uq.queryString += fmt.Sprintf(", %s", s.ToString())
		}
	}

	for i, w := range uq.where {
		if i == 0 {
			uq.queryString += fmt.Sprintf("WHERE %s", w.ToString())
		} else {
			uq.queryString += fmt.Sprintf("AND %s", w.ToString())
		}
	}

	uq.queryString = strings.ReplaceAll(uq.queryString, "$jk0__table_name$", uq.origin.TableExpr())

	return gc.conn.Exec(uq.queryString)
}
