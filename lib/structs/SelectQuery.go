package structs

import (
	"fmt"
	"gormy/lib/engine"
	"gormy/lib/joins"
	"gormy/lib/sqlparser"
	"gormy/lib/types"
	"strings"
)

type SelectQuery[T any] struct {
	Query[T]
	selected        []types.Column
	activeRelations []ActiveRelation
}

func (query *SelectQuery[T]) Column(columnName string) *SelectQuery[T] {

	column, err := query.columnByName(columnName)

	if err != nil {
		panic(err)
	}

	query.selected = append(query.selected, column)

	return query
}

func (query *SelectQuery[T]) Where(expr string, columnName string, value string) *SelectQuery[T] {

	cleanExpr := strings.ReplaceAll(expr, "?", "%s")

	query.queryString = query.queryString + fmt.Sprintf("WHERE "+cleanExpr, columnName, value)

	return query
}

func (query *SelectQuery[T]) Relation(relationName string, joinName string) *SelectQuery[T] {

	relation, err := query.relationByName(relationName)

	if err != nil {
		println(err.Error())
	}

	query.origin.OnJoin()

	join, err := joins.Joins.ByName(joinName)

	activeRelation := ActiveRelation{
		Relation: relation,
		Join:     join(relation.JoinKey, relation.Name, query.tableName, relation.Columns, relation.TableName),
	}

	query.activeRelations = append(query.activeRelations, activeRelation)

	query.queryString += activeRelation.Join.JoinExpr(relation)
	// query.queryString += fmt.Sprintf("%s JOIN $%s__table_name$ %s ON %s.%s = %s.%s", relation.How, relation.JoinKey, relation.JoinKey, "jk0", relation.Key, relation.JoinKey, relation.ForeignKey)

	return query
}

func (query *SelectQuery[T]) Exec() []T {
	queryString := query.queryString

	columnExpr := ""

	if len(query.selected) != 0 {
		for _, column := range query.selected {
			if column == query.selected[len(query.selected)-1] {
				columnExpr += fmt.Sprintf("%s.%s as %s__%s", "jk0", column.Name, "jk0", column.StructName)
			} else {
				columnExpr += fmt.Sprintf("%s.%s as %s__%s", "jk0", column.Name, "jk0", column.StructName) + ", "
			}
		}
	} else {
		columnExpr = query.origin.ColumnExpr()

		for _, relation := range query.activeRelations {
			columnExpr += ", " + relation.Join.ColumnExpr()
		}
	}

	queryString = strings.ReplaceAll(queryString, "$columns$", columnExpr)

	queryString = strings.ReplaceAll(queryString, "$jk0__table_name$", query.origin.TableExpr())

	for _, relation := range query.activeRelations {
		queryString = strings.ReplaceAll(queryString, fmt.Sprintf("$%s__table_name$", relation.Join.JoinKey()), relation.Join.TableExpr())
	}

	queryString += "\r\nORDER BY jk0__join_row"

	rows, err := engine.DB().Query(queryString)

	if err != nil {
		println(err.Error())
	}

	joins := []joins.Join{}

	for _, ar := range query.activeRelations {
		joins = append(joins, ar.Join)
	}

	sqlParser := sqlparser.NewSqlParser(query.Rows, joins, *rows)

	sqlParser.Parse(&query.Rows)

	// keys := make([]string, 0, len(query.relations))
	// names := make([]string, 0, len(query.relations))
	// // for k := range query.relationMap {
	// // 	keys = append(keys, k)
	// // 	names = append(names, query.relationMap[k].Name)
	// // }

	// jsonString, err := utils.RowsToJson(*rows, keys, names, map[string]string{"jk1": "onetomany"})

	// println(jsonString)

	// json.Unmarshal([]byte(jsonString), &query.Rows)

	// if err != nil {
	// 	println(err.Error())
	// }

	return query.Rows
}
