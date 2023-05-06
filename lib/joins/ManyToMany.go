package joins

import (
	"database/sql"
	"fmt"
	"gormy/lib/types"
)

type manyToMany struct {
	joinkey        string
	values         map[int][]map[string]interface{}
	joinName       string
	columns        []types.Column
	derivedColumns []string
	intermediary   string
	joinsTo        string
	tableExpr      string
}

func (manytomany manyToMany) Columns() []types.Column {
	return manytomany.columns
}

func (manytomany manyToMany) ColumnExpr() string {
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

func (manytomany manyToMany) TableExpr() string {
	return manytomany.tableExpr
}

func (manyToMany manyToMany) JoinExpr(relation types.Relation) string {
	str_out := ""
	if _, ok := relation.TagData["intermediary"]; ok {
		str_out += fmt.Sprintf("%s JOIN %s %s_i ON %s.%s = %s_i.%s\r\n", relation.How, relation.TagData["intermediary"], relation.JoinKey, "jk0", relation.Key, relation.JoinKey, relation.Key)
		str_out += fmt.Sprintf("%s JOIN $%s__table_name$ %s ON %s_i.%s = %s.%s", relation.How, relation.JoinKey, relation.JoinKey, relation.JoinKey, relation.ForeignKey, relation.JoinKey, relation.ForeignKey)
	}

	return str_out
}

func (manytomany manyToMany) OnJoin() {
	manytomany.tableExpr = fmt.Sprintf("(select *, row_number () over() as %s__join_row from %s)", manytomany.joinkey, manytomany.tableExpr)
}

func (manytomany manyToMany) Parser(rowNumber int, repeatRowNumber int, key string, name string, column *sql.ColumnType, sqlType interface{}) {

	if key == manytomany.joinkey {
		if _, ok := manytomany.values[rowNumber]; !ok {
			manytomany.values[rowNumber] = []map[string]interface{}{}
		}

		if len(manytomany.values[rowNumber]) <= repeatRowNumber {
			manytomany.values[rowNumber] = append(manytomany.values[rowNumber], map[string]interface{}{})
		}

		if z, ok := (sqlType).(*sql.NullBool); ok {
			manytomany.values[rowNumber][repeatRowNumber][name] = z.Bool
			return
		}

		if z, ok := (sqlType).(*sql.NullString); ok {
			manytomany.values[rowNumber][repeatRowNumber][name] = z.String
			return
		}

		if z, ok := (sqlType).(*sql.NullInt64); ok {
			manytomany.values[rowNumber][repeatRowNumber][name] = z.Int64
			return
		}

		if z, ok := (sqlType).(*sql.NullFloat64); ok {
			manytomany.values[rowNumber][repeatRowNumber][name] = z.Float64
			return
		}

		if z, ok := (sqlType).(*sql.NullInt32); ok {
			manytomany.values[rowNumber][repeatRowNumber][name] = z.Int32
			return
		}

		manytomany.values[rowNumber][repeatRowNumber][name] = sqlType
		return

	}
}

func (manytomany manyToMany) Values(rowNumber int) interface{} {
	return manytomany.values[rowNumber]
}

func (manytomany manyToMany) JoinName() string {
	return manytomany.joinName
}

func (manytomany manyToMany) JoinKey() string {
	return manytomany.joinkey
}

func ManyToMany(joinkey string, joinName string, joinsTo string, columns []types.Column, tableExpr string) Join {
	return manyToMany{
		joinkey:   joinkey,
		values:    map[int][]map[string]interface{}{},
		joinName:  joinName,
		columns:   columns,
		joinsTo:   joinsTo,
		tableExpr: tableExpr,
	}
}
