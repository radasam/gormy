package joins

import (
	"database/sql"
	"fmt"
	"gormy/lib/types"
)

type oneToMany struct {
	joinkey        string
	values         map[int][]map[string]interface{}
	joinName       string
	columns        []types.Column
	derivedColumns []string
	joinsTo        string
	tableExpr      string
}

func (onetomany oneToMany) Columns() []types.Column {
	return onetomany.columns
}

func (onetomany oneToMany) ColumnExpr() string {
	columnExpr := ""

	for _, c := range onetomany.derivedColumns {
		columnExpr += fmt.Sprintf("%s, ", c)
	}

	for i, c := range onetomany.columns {
		if i == len(onetomany.columns)-1 {
			columnExpr += fmt.Sprintf("%s.%s as %s__%s", onetomany.joinkey, c.Name, onetomany.joinkey, c.StructName)
		} else {
			columnExpr += fmt.Sprintf("%s.%s as %s__%s,", onetomany.joinkey, c.Name, onetomany.joinkey, c.StructName)
		}
	}

	return columnExpr
}

func (onetomany oneToMany) TableExpr() string {
	return onetomany.tableExpr
}

func (onetomany oneToMany) JoinExpr(relation types.Relation) string {
	return ""
}

func (onetomany oneToMany) OnJoin() {
	onetomany.tableExpr = fmt.Sprintf("(select *, row_number () over() as %s__join_row from %s)", onetomany.joinkey, onetomany.tableExpr)
}

func (onetomany oneToMany) Parser(rowNumber int, repeatRowNumber int, key string, name string, column *sql.ColumnType, sqlType interface{}) {

	if key == onetomany.joinkey {
		if _, ok := onetomany.values[rowNumber]; !ok {
			onetomany.values[rowNumber] = []map[string]interface{}{}
		}

		if len(onetomany.values[rowNumber]) <= repeatRowNumber {
			onetomany.values[rowNumber] = append(onetomany.values[rowNumber], map[string]interface{}{})
		}

		if z, ok := (sqlType).(*sql.NullBool); ok {
			onetomany.values[rowNumber][repeatRowNumber][name] = z.Bool
			return
		}

		if z, ok := (sqlType).(*sql.NullString); ok {
			onetomany.values[rowNumber][repeatRowNumber][name] = z.String
			return
		}

		if z, ok := (sqlType).(*sql.NullInt64); ok {
			onetomany.values[rowNumber][repeatRowNumber][name] = z.Int64
			return
		}

		if z, ok := (sqlType).(*sql.NullFloat64); ok {
			onetomany.values[rowNumber][repeatRowNumber][name] = z.Float64
			return
		}

		if z, ok := (sqlType).(*sql.NullInt32); ok {
			onetomany.values[rowNumber][repeatRowNumber][name] = z.Int32
			return
		}

		onetomany.values[rowNumber][repeatRowNumber][name] = sqlType
		return

	}
}

func (onetomany oneToMany) onJoin() {
	onetomany.tableExpr = fmt.Sprintf("(select *, row_number () over() as %s__join_row from %s)", onetomany.joinkey, onetomany.tableExpr)
}

func (onetomany oneToMany) Values(rowNumber int) interface{} {
	return onetomany.values[rowNumber]
}

func (onetomany oneToMany) JoinName() string {
	return onetomany.joinName
}

func (onetomany oneToMany) JoinKey() string {
	return onetomany.joinkey
}

func OneToMany(joinkey string, joinName string, joinsTo string, columns []types.Column, tableExpr string) Join {
	return oneToMany{
		joinkey:   joinkey,
		values:    map[int][]map[string]interface{}{},
		joinName:  joinName,
		columns:   columns,
		joinsTo:   joinsTo,
		tableExpr: tableExpr,
	}
}
