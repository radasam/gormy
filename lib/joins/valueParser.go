package joins

import (
	"database/sql"
	"fmt"
	"strings"
)

type sqlValueParser struct {
	joinkey   string
	joins     []Join
	values    map[string]map[string]interface{}
	rowNumber int
}

func (sqlvalueparser *sqlValueParser) Parse(parentRow string, key string, name string, column *sql.ColumnType, sqlType interface{}) {
	if strings.HasPrefix(key, sqlvalueparser.joinkey) {
		if key == sqlvalueparser.joinkey {
			if name == "join_row" {
				if z, ok := (sqlType).(*sql.NullInt64); ok {
					sqlvalueparser.rowNumber = int(z.Int64) - 1
					return
				}
				if z, ok := (sqlType).(*sql.NullInt32); ok {
					sqlvalueparser.rowNumber = int(z.Int32) - 1
					return
				}
			} else {
				if _, ok := sqlvalueparser.values[parentRow]; !ok {
					sqlvalueparser.values[parentRow] = map[string]interface{}{}
				}

				if z, ok := (sqlType).(*sql.NullBool); ok {
					sqlvalueparser.values[parentRow][name] = z.Bool
					return
				}

				if z, ok := (sqlType).(*sql.NullString); ok {
					sqlvalueparser.values[parentRow][name] = z.String
					return
				}

				if z, ok := (sqlType).(*sql.NullInt64); ok {
					sqlvalueparser.values[parentRow][name] = z.Int64
					return
				}

				if z, ok := (sqlType).(*sql.NullFloat64); ok {
					sqlvalueparser.values[parentRow][name] = z.Float64
					return
				}

				if z, ok := (sqlType).(*sql.NullInt32); ok {
					sqlvalueparser.values[parentRow][name] = z.Int32
					return
				}

				sqlvalueparser.values[parentRow][name] = sqlType
				return
			}
		} else {
			for _, join := range sqlvalueparser.joins {
				join.Parse(fmt.Sprintf("%s_%d", parentRow, sqlvalueparser.rowNumber), key, name, column, sqlType)
				sqlvalueparser.values[parentRow][join.JoinName()] = join.Row(fmt.Sprintf("%s_%d", parentRow, sqlvalueparser.rowNumber))
			}
		}
	}
	return
}

func (sqlvalueparser *sqlValueParser) Row(parentRow string) interface{} {
	return sqlvalueparser.values[parentRow]
}

func (sqlValueParser *sqlValueParser) Values() interface{} {
	return sqlValueParser.values
}

func (sqlvalueparser *sqlValueParser) OnJoin(join Join) {
	sqlvalueparser.joins = append(sqlvalueparser.joins, join)
}

func NewValueParser(joinKey string) sqlValueParser {
	return sqlValueParser{
		joinkey: joinKey, joins: []Join{}, values: map[string]map[string]interface{}{}, rowNumber: -1,
	}
}
