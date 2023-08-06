package gormy

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/radasam/gormy/internal/driver"
)

type joinValueParser struct {
	joinkey   string
	joins     []Join
	values    map[string]map[string]interface{}
	rowNumber int
}

func (jvp *joinValueParser) Parse(parentRow string, key string, name string, column driver.ColumnType, sqlType interface{}) {
	if strings.HasPrefix(key, jvp.joinkey) {
		if key == jvp.joinkey {
			if name == "join_row" {
				if z, ok := (sqlType).(*sql.NullInt64); ok {
					jvp.rowNumber = int(z.Int64) - 1
					return
				}
				if z, ok := (sqlType).(*sql.NullInt32); ok {
					jvp.rowNumber = int(z.Int32) - 1
					return
				}
			} else {
				if _, ok := jvp.values[parentRow]; !ok {
					jvp.values[parentRow] = map[string]interface{}{}
				}

				if z, ok := (sqlType).(*sql.NullBool); ok {
					jvp.values[parentRow][name] = z.Bool
					return
				}

				if z, ok := (sqlType).(*sql.NullString); ok {
					jvp.values[parentRow][name] = z.String
					return
				}

				if z, ok := (sqlType).(*sql.NullInt64); ok {
					jvp.values[parentRow][name] = z.Int64
					return
				}

				if z, ok := (sqlType).(*sql.NullFloat64); ok {
					jvp.values[parentRow][name] = z.Float64
					return
				}

				if z, ok := (sqlType).(*sql.NullInt32); ok {
					jvp.values[parentRow][name] = z.Int32
					return
				}

				jvp.values[parentRow][name] = sqlType
				return
			}
		} else {
			for _, join := range jvp.joins {
				join.Parse(fmt.Sprintf("%s_%d", parentRow, jvp.rowNumber), key, name, column, sqlType)
				jvp.values[parentRow][join.JoinName()] = join.Row(fmt.Sprintf("%s_%d", parentRow, jvp.rowNumber))
			}
		}
	}
}

func (jvp *joinValueParser) Row(parentRow string) interface{} {
	return jvp.values[parentRow]
}

func (joinValueParser *joinValueParser) Values() interface{} {
	return joinValueParser.values
}

func (jvp *joinValueParser) OnJoin(join Join) {
	jvp.joins = append(jvp.joins, join)
}

func newJoinValueParser(joinKey string) joinValueParser {
	return joinValueParser{
		joinkey: joinKey, joins: []Join{}, values: map[string]map[string]interface{}{}, rowNumber: -1,
	}
}
