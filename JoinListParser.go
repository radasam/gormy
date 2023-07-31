package gormy

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
)

type joinListParser struct {
	joinkey             string
	joins               []Join
	values              map[string][]map[string]interface{}
	rowNumber           int
	repeatRowNumber     int
	lastParentRowNumber string
	lastRowNumber       int
	id                  int
}

func (jlp *joinListParser) Parse(parentRow string, key string, name string, column *sql.ColumnType, sqlType interface{}) {

	if strings.HasPrefix(key, jlp.joinkey) {
		if key == jlp.joinkey {

			if name == "join_row" {

				if z, ok := (sqlType).(*sql.NullInt64); ok {
					jlp.rowNumber = int(z.Int64) - 1
					return
				}
				if z, ok := (sqlType).(*sql.NullInt32); ok {
					jlp.rowNumber = int(z.Int32) - 1
					return
				}

				if jlp.lastRowNumber == jlp.rowNumber {
					if parentRow == jlp.lastParentRowNumber {
						jlp.repeatRowNumber += 1
					} else {
						jlp.repeatRowNumber = 0
					}
				}

				jlp.lastParentRowNumber = parentRow

				jlp.lastRowNumber = jlp.rowNumber
			} else {
				if _, ok := jlp.values[parentRow]; !ok {
					jlp.values[parentRow] = []map[string]interface{}{}
				}

				if len(jlp.values[parentRow]) == jlp.rowNumber {
					jlp.values[parentRow] = append(jlp.values[parentRow], map[string]interface{}{})
				}

				if z, ok := (sqlType).(*sql.NullBool); ok {
					jlp.values[parentRow][jlp.rowNumber][name] = z.Bool
					return
				}

				if z, ok := (sqlType).(*sql.NullString); ok {
					jlp.values[parentRow][jlp.rowNumber][name] = z.String
					return
				}

				if z, ok := (sqlType).(*sql.NullInt64); ok {
					jlp.values[parentRow][jlp.rowNumber][name] = z.Int64
					return
				}

				if z, ok := (sqlType).(*sql.NullFloat64); ok {
					jlp.values[parentRow][jlp.rowNumber][name] = z.Float64
					return
				}

				if z, ok := (sqlType).(*sql.NullInt32); ok {
					jlp.values[parentRow][jlp.rowNumber][name] = z.Int32
					return
				}

				jlp.values[parentRow][jlp.rowNumber][name] = sqlType
				return

			}

		} else {
			for _, join := range jlp.joins {
				join.Parse(fmt.Sprintf("%s_%d", parentRow, jlp.rowNumber), key, name, column, sqlType)
				jlp.values[parentRow][jlp.rowNumber][join.JoinName()] = join.Row(fmt.Sprintf("%s_%d", parentRow, jlp.rowNumber))
			}
		}
	}

	return
}

func (jlp *joinListParser) Row(parentRow string) interface{} {
	return jlp.values[parentRow]
}

func (jlp *joinListParser) Values() interface{} {
	return jlp.values
}
func (jlp *joinListParser) OnJoin(join Join) {
	jlp.joins = append(jlp.joins, join)
}

func newJoinListParser(joinKey string) *joinListParser {
	return &joinListParser{
		joinkey: joinKey, joins: []Join{}, values: map[string][]map[string]interface{}{}, rowNumber: 0, repeatRowNumber: 0, lastRowNumber: -1, lastParentRowNumber: "", id: rand.Int(),
	}
}
