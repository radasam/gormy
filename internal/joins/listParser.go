package joins

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
)

type sqlListParser struct {
	joinkey             string
	joins               []Join
	values              map[string][]map[string]interface{}
	rowNumber           int
	repeatRowNumber     int
	lastParentRowNumber string
	lastRowNumber       int
	id                  int
}

func (sqllistparser *sqlListParser) Parse(parentRow string, key string, name string, column *sql.ColumnType, sqlType interface{}) {

	if strings.HasPrefix(key, sqllistparser.joinkey) {
		if key == sqllistparser.joinkey {

			if name == "join_row" {

				if z, ok := (sqlType).(*sql.NullInt64); ok {
					sqllistparser.rowNumber = int(z.Int64) - 1
					return
				}
				if z, ok := (sqlType).(*sql.NullInt32); ok {
					sqllistparser.rowNumber = int(z.Int32) - 1
					return
				}

				if sqllistparser.lastRowNumber == sqllistparser.rowNumber {
					if parentRow == sqllistparser.lastParentRowNumber {
						sqllistparser.repeatRowNumber += 1
					} else {
						sqllistparser.repeatRowNumber = 0
					}
				}

				sqllistparser.lastParentRowNumber = parentRow

				sqllistparser.lastRowNumber = sqllistparser.rowNumber
			} else {
				if _, ok := sqllistparser.values[parentRow]; !ok {
					sqllistparser.values[parentRow] = []map[string]interface{}{}
				}

				if len(sqllistparser.values[parentRow]) == sqllistparser.rowNumber {
					sqllistparser.values[parentRow] = append(sqllistparser.values[parentRow], map[string]interface{}{})
				}

				if z, ok := (sqlType).(*sql.NullBool); ok {
					sqllistparser.values[parentRow][sqllistparser.rowNumber][name] = z.Bool
					return
				}

				if z, ok := (sqlType).(*sql.NullString); ok {
					sqllistparser.values[parentRow][sqllistparser.rowNumber][name] = z.String
					return
				}

				if z, ok := (sqlType).(*sql.NullInt64); ok {
					sqllistparser.values[parentRow][sqllistparser.rowNumber][name] = z.Int64
					return
				}

				if z, ok := (sqlType).(*sql.NullFloat64); ok {
					sqllistparser.values[parentRow][sqllistparser.rowNumber][name] = z.Float64
					return
				}

				if z, ok := (sqlType).(*sql.NullInt32); ok {
					sqllistparser.values[parentRow][sqllistparser.rowNumber][name] = z.Int32
					return
				}

				sqllistparser.values[parentRow][sqllistparser.rowNumber][name] = sqlType
				return

			}

		} else {
			for _, join := range sqllistparser.joins {
				join.Parse(fmt.Sprintf("%s_%d", parentRow, sqllistparser.rowNumber), key, name, column, sqlType)
				sqllistparser.values[parentRow][sqllistparser.rowNumber][join.JoinName()] = join.Row(fmt.Sprintf("%s_%d", parentRow, sqllistparser.rowNumber))
			}
		}
	}

	return
}

func (sqllistparser *sqlListParser) Row(parentRow string) interface{} {
	return sqllistparser.values[parentRow]
}

func (sqllistparser *sqlListParser) Values() interface{} {
	return sqllistparser.values
}
func (sqllistparser *sqlListParser) OnJoin(join Join) {
	sqllistparser.joins = append(sqllistparser.joins, join)
}

func NewListParser(joinKey string) *sqlListParser {
	return &sqlListParser{
		joinkey: joinKey, joins: []Join{}, values: map[string][]map[string]interface{}{}, rowNumber: 0, repeatRowNumber: 0, lastRowNumber: -1, lastParentRowNumber: "", id: rand.Int(),
	}
}
