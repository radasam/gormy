package joins

import (
	"database/sql"
	"fmt"
	"gormy/lib/types"
)

type oneToOne struct {
	joinkey        string
	values         map[int]map[string]interface{}
	joinName       string
	columns        []types.Column
	derivedColumns []string
	joinsTo        string
	tableExpr      string
}

func (onetoone oneToOne) Columns() []types.Column {
	return onetoone.columns
}

func (onetoone oneToOne) ColumnExpr() string {
	columnExpr := ""

	for _, c := range onetoone.derivedColumns {
		columnExpr += fmt.Sprintf("%s, ", c)
	}

	for i, c := range onetoone.columns {
		if i == len(onetoone.columns)-1 {
			columnExpr += fmt.Sprintf("%s.%s as %s__%s", onetoone.joinkey, c.Name, onetoone.joinkey, c.StructName)
		} else {
			columnExpr += fmt.Sprintf("%s.%s as %s__%s,", onetoone.joinkey, c.Name, onetoone.joinkey, c.StructName)
		}
	}

	return columnExpr
}

func (onetoone oneToOne) TableExpr() string {
	return onetoone.tableExpr
}

func (onetoone oneToOne) JoinExpr(relation types.Relation) string {
	return ""
}

func (onetoone oneToOne) OnJoin() {
}

func (onetoone oneToOne) Parser(rowNumber int, repeatRowNumber int, key string, name string, column *sql.ColumnType, sqlType interface{}) {

	if key == onetoone.joinkey {
		if _, ok := onetoone.values[rowNumber]; !ok {
			onetoone.values[rowNumber] = map[string]interface{}{}
		}
		if z, ok := (sqlType).(*sql.NullBool); ok {
			onetoone.values[rowNumber][name] = z.Bool
			return
		}

		if z, ok := (sqlType).(*sql.NullString); ok {
			onetoone.values[rowNumber][name] = z.String
			return
		}

		if z, ok := (sqlType).(*sql.NullInt64); ok {
			onetoone.values[rowNumber][name] = z.Int64
			return
		}

		if z, ok := (sqlType).(*sql.NullFloat64); ok {
			onetoone.values[rowNumber][name] = z.Float64
			return
		}

		if z, ok := (sqlType).(*sql.NullInt32); ok {
			onetoone.values[rowNumber][name] = z.Int32
			return
		}

		onetoone.values[rowNumber][name] = sqlType

	}
}

func (onetoone oneToOne) onJoin() {
	onetoone.tableExpr = fmt.Sprintf("(select *, row_number () over() as %s__join_row from %s)", onetoone.joinkey, onetoone.tableExpr)
}

func (onetoone oneToOne) Values(rowNumber int) interface{} {
	return onetoone.values
}

func (onetoone oneToOne) JoinName() string {
	return onetoone.joinName
}

func (onetoone oneToOne) JoinKey() string {
	return onetoone.joinkey
}

func OneToOne(joinkey string, joinName string, joinsTo string, columns []types.Column, tableExpr string) Join {
	return oneToOne{
		joinkey:   joinkey,
		values:    map[int]map[string]interface{}{},
		joinName:  joinName,
		columns:   columns,
		joinsTo:   joinsTo,
		tableExpr: tableExpr,
	}
}
