package joins

import (
	"database/sql"
	"fmt"
	"gormy/lib/types"
)

type origin struct {
	joinkey        string
	values         map[int]map[string]interface{}
	joinName       string
	columns        []types.Column
	derivedColumns []string
	tableExpr      string
	parser         sqlValueParser
}

func (_origin origin) Columns() []types.Column {
	return _origin.columns
}

func (_origin origin) ColumnExpr() string {
	columnExpr := ""

	for _, c := range _origin.derivedColumns {
		columnExpr += fmt.Sprintf("%s, ", c)
	}

	for i, c := range _origin.columns {
		if i == len(_origin.columns)-1 {
			columnExpr += fmt.Sprintf("%s.%s as %s__%s", _origin.joinkey, c.Name, _origin.joinkey, c.StructName)
		} else {
			columnExpr += fmt.Sprintf("%s.%s as %s__%s,", _origin.joinkey, c.Name, _origin.joinkey, c.StructName)
		}
	}

	return columnExpr
}

func (_origin *origin) TableExpr() string {
	return _origin.tableExpr
}

func (_origin *origin) JoinExpr(originKey string, relation types.Relation) string {
	return ""
}
func (_origin *origin) Parse(parentRow string, key string, name string, column *sql.ColumnType, sqlType interface{}) {
	_origin.parser.Parse(parentRow, key, name, column, sqlType)
}
func (_origin *origin) OnJoin(join Join) {
	_origin.tableExpr = fmt.Sprintf("(select *, row_number () over() as jk0__join_row from %s)", _origin.tableExpr)
	_origin.derivedColumns = append(_origin.derivedColumns, "jk0__join_row")
	_origin.parser.OnJoin(join)
}

func (_origin origin) Row(parentRow string) interface{} {
	return _origin.parser.Row(parentRow)
}

func (_origin origin) Values() interface{} {
	mapValues := _origin.parser.Values()
	values := []interface{}{}

	if _, ok := mapValues.(map[string]map[string]interface{}); ok {
		for _, v := range mapValues.(map[string]map[string]interface{}) {
			values = append(values, v)
		}
	}

	return values
}

func (_origin origin) JoinName() string {
	return _origin.joinName
}

func (_origin origin) JoinKey() string {
	return _origin.joinkey
}

func Origin(joinName string, columns []types.Column, tableExpr string) Join {
	return &origin{
		joinkey:   "jk0",
		values:    map[int]map[string]interface{}{},
		joinName:  joinName,
		columns:   columns,
		tableExpr: tableExpr,
		parser:    NewValueParser("jk0"),
	}
}
