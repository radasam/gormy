package gormy

import (
	"database/sql"
	"fmt"
)

type oneToOne struct {
	joinkey        string
	values         map[int]map[string]interface{}
	joinName       string
	columns        []Column
	derivedColumns []string
	joinsTo        string
	tableExpr      string
	parser         joinValueParser
	parentJoinRow  string
	hasJoin        bool
}

func (onetoone oneToOne) Columns() []Column {
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

func (onetoone oneToOne) JoinExpr(originKey string, relation Relation) string {
	return fmt.Sprintf("%s JOIN $%s__table_name$ %s ON %s.%s = %s.%s\r\n", relation.How, relation.JoinKey, relation.JoinKey, originKey, relation.Key, relation.JoinKey, relation.ForeignKey)
}

func (onetoone oneToOne) OnJoin(join Join) {
	onetoone.hasJoin = false
	onetoone.tableExpr = fmt.Sprintf("(select *, row_number () over(partition by %s) as %s__join_row from %s)", onetoone.parentJoinRow, onetoone.joinkey, onetoone.tableExpr)
	onetoone.parser.OnJoin(join)
}

func (onetoone oneToOne) HasJoin() bool {
	return onetoone.hasJoin
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

func (onetoone oneToOne) Parse(parentRow string, key string, name string, column *sql.ColumnType, sqlType interface{}) {
	onetoone.parser.Parse(parentRow, key, name, column, sqlType)
}

func (onetoone oneToOne) Row(parentRow string) interface{} {
	return onetoone.parser.Row(parentRow)
}

func (onetoone oneToOne) Values() interface{} {
	return onetoone.parser.Values()
}

func (onetoone oneToOne) JoinName() string {
	return onetoone.joinName
}

func (onetoone oneToOne) JoinKey() string {
	return onetoone.joinkey
}

func OneToOne(joinkey string, joinName string, joinsTo string, columns []Column, tableExpr string, parentJoinRow string) Join {
	return oneToOne{
		joinkey:       joinkey,
		values:        map[int]map[string]interface{}{},
		joinName:      joinName,
		columns:       columns,
		joinsTo:       joinsTo,
		tableExpr:     tableExpr,
		parser:        newJoinValueParser(joinkey),
		parentJoinRow: parentJoinRow,
	}
}
