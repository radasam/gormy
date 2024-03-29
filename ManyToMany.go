package gormy

import (
	"fmt"

	"github.com/radasam/gormy/internal/driver"
)

type manyToMany struct {
	joinkey        string
	values         map[int][]map[string]interface{}
	joinName       string
	columns        []Column
	derivedColumns []string
	intermediary   string
	joinsTo        string
	tableExpr      string
	parser         *joinListParser
	parentJoinRow  string
	hasJoin        bool
}

func (manytomany *manyToMany) Columns() []Column {
	return manytomany.columns
}

func (manytomany *manyToMany) ColumnExpr() string {
	columnExpr := ""

	for _, c := range manytomany.derivedColumns {
		columnExpr += fmt.Sprintf("%s, ", c)
	}

	for i, c := range manytomany.columns {
		if i == len(manytomany.columns)-1 {
			columnExpr += fmt.Sprintf("%s.%s as %s__%s", manytomany.joinkey, c.Name, manytomany.joinkey, c.StructName)
		} else {
			columnExpr += fmt.Sprintf("%s.%s as %s__%s,", manytomany.joinkey, c.Name, manytomany.joinkey, c.StructName)
		}
	}

	return columnExpr
}

func (manytomany *manyToMany) TableExpr() string {
	return manytomany.tableExpr
}

func (manytomany *manyToMany) JoinExpr(originKey string, relation Relation) string {
	str_out := ""
	if _, ok := relation.TagData["intermediary"]; ok {
		str_out += fmt.Sprintf("%s JOIN %s %s_i ON %s.%s = %s_i.%s\r\n", relation.How, relation.TagData["intermediary"], relation.JoinKey, originKey, relation.Key, relation.JoinKey, relation.Key)
		str_out += fmt.Sprintf("%s JOIN $%s__table_name$ %s ON %s_i.%s = %s.%s\r\n", relation.How, relation.JoinKey, relation.JoinKey, relation.JoinKey, relation.ForeignKey, relation.JoinKey, relation.ForeignKey)
	}

	return str_out
}

func (manytomany *manyToMany) OnJoin(join Join) {
	manytomany.hasJoin = true
	manytomany.tableExpr = fmt.Sprintf("(select *, row_number () over(partition by %s) as %s__join_row from %s)", manytomany.parentJoinRow, manytomany.joinkey, manytomany.tableExpr)
	manytomany.derivedColumns = append(manytomany.derivedColumns, fmt.Sprintf("%s__join_row", manytomany.joinkey))
	manytomany.parser.OnJoin(join)
}

func (manytomany *manyToMany) HasJoin() bool {
	return manytomany.hasJoin
}

func (manytomany *manyToMany) Parse(parentRow string, key string, name string, column driver.ColumnType, sqlType interface{}) {
	manytomany.parser.Parse(parentRow, key, name, column, sqlType)
}

func (manytomany *manyToMany) Row(parentRow string) interface{} {
	return manytomany.parser.Row(parentRow)
}

func (manytomany *manyToMany) Values() interface{} {
	return manytomany.parser.Values()
}

func (manytomany *manyToMany) JoinName() string {
	return manytomany.joinName
}

func (manytomany *manyToMany) JoinKey() string {
	return manytomany.joinkey
}

func ManyToMany(joinkey string, joinName string, joinsTo string, columns []Column, tableExpr string, parentJoinRow string) Join {
	return &manyToMany{
		joinkey:       joinkey,
		values:        map[int][]map[string]interface{}{},
		joinName:      joinName,
		columns:       columns,
		joinsTo:       joinsTo,
		tableExpr:     tableExpr,
		parser:        newJoinListParser(joinkey),
		parentJoinRow: parentJoinRow,
	}
}
