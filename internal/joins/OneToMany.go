package joins

import (
	"database/sql"
	"fmt"

	"github.com/radasam/gormy/internal/types"
)

type oneToMany struct {
	joinkey        string
	values         map[int][]map[string]interface{}
	joinName       string
	columns        []types.Column
	derivedColumns []string
	joinsTo        string
	tableExpr      string
	joins          []Join
	parser         SqlParser
	parentJoinRow  string
}

func (onetomany *oneToMany) Columns() []types.Column {
	return onetomany.columns
}

func (onetomany *oneToMany) ColumnExpr() string {
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

func (onetomany *oneToMany) TableExpr() string {
	return onetomany.tableExpr
}

func (onetomany *oneToMany) JoinExpr(originKey string, relation types.Relation) string {
	return fmt.Sprintf("%s JOIN $%s__table_name$ %s ON %s.%s = %s.%s\r\n", relation.How, relation.JoinKey, relation.JoinKey, originKey, relation.Key, relation.JoinKey, relation.ForeignKey)
}

func (onetomany *oneToMany) OnJoin(join Join) {
	onetomany.parser.OnJoin(join)
}

func (onetomany *oneToMany) Parse(parentRow string, key string, name string, column *sql.ColumnType, sqlType interface{}) {
	onetomany.parser.Parse(parentRow, key, name, column, sqlType)
}

func (onetomany *oneToMany) Row(parentRow string) interface{} {
	return onetomany.parser.Row(parentRow)
}

func (onetomany *oneToMany) Values() interface{} {
	return onetomany.parser.Values()
}

func (onetomany *oneToMany) JoinName() string {
	return onetomany.joinName
}

func (onetomany *oneToMany) JoinKey() string {
	return onetomany.joinkey
}

func (onetomany *oneToMany) AddJoin(join Join) {
	onetomany.joins = append(onetomany.joins, join)
}

func OneToMany(joinkey string, joinName string, joinsTo string, columns []types.Column, tableExpr string, parentJoinRow string) Join {
	tableExpr = fmt.Sprintf("(select *, row_number () over(partition by %s) as %s__join_row from %s)", parentJoinRow, joinkey, tableExpr)
	return &oneToMany{
		joinkey:        joinkey,
		values:         map[int][]map[string]interface{}{},
		joinName:       joinName,
		columns:        columns,
		joinsTo:        joinsTo,
		tableExpr:      tableExpr,
		parser:         NewListParser(joinkey),
		parentJoinRow:  parentJoinRow,
		derivedColumns: []string{fmt.Sprintf("%s__join_row", joinkey)},
	}
}
