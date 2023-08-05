package gormy

import (
	"fmt"
	"strings"
)

type Statement struct {
	expr       string
	columnName string
	value      string
}

func (ss Statement) ToString() string {
	cleanExpr := strings.ReplaceAll(ss.expr, "?", "%s")

	return fmt.Sprintf(cleanExpr+"\r\n", ss.columnName, ss.value)
}
