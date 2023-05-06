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

func (_origin *origin) JoinExpr(relation types.Relation) string {
	return ""
}

func (_origin origin) Parser(rowNumber int, repeatRowNumber int, key string, name string, column *sql.ColumnType, sqlType interface{}) {

	if key == _origin.joinkey {
		if _, ok := _origin.values[rowNumber]; !ok {
			_origin.values[rowNumber] = map[string]interface{}{}
		}
		if z, ok := (sqlType).(*sql.NullBool); ok {
			_origin.values[rowNumber][name] = z.Bool
			return
		}

		if z, ok := (sqlType).(*sql.NullString); ok {
			_origin.values[rowNumber][name] = z.String
			return
		}

		if z, ok := (sqlType).(*sql.NullInt64); ok {
			_origin.values[rowNumber][name] = z.Int64
			return
		}

		if z, ok := (sqlType).(*sql.NullFloat64); ok {
			_origin.values[rowNumber][name] = z.Float64
			return
		}

		if z, ok := (sqlType).(*sql.NullInt32); ok {
			_origin.values[rowNumber][name] = z.Int32
			return
		}

		_origin.values[rowNumber][name] = sqlType

	}
}

func (_origin *origin) OnJoin() {
	_origin.tableExpr = fmt.Sprintf("(select *, row_number () over() as jk0__join_row from %s)", _origin.tableExpr)
	_origin.derivedColumns = append(_origin.derivedColumns, "jk0__join_row")
}

func (_origin origin) Values(rowNumber int) interface{} {
	return _origin.values
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
	}
}
