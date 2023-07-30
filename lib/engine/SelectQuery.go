package engine

import (
	"fmt"
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

func (query *SelectQuery[T]) relationByName(relationName string, joinName string) (string, types.Relation, joins.Join, error) {

	joinType, err := joins.Joins.ByName(joinName)

	if err != nil {
		return "", types.Relation{}, nil, fmt.Errorf("Relation doesnt exist")
	}

	for _, relation := range query.relations {
		if relation.Name == relationName {
			join := joinType(relation.JoinKey, relation.Name, query.tableName, relation.Columns, relation.TableName, relation.ForeignKey)
			query.origin.OnJoin(join)
			return query.origin.JoinKey(), relation, join, nil
		}
	}

	for _, ar := range query.activeRelations {
		relation, err := ar.Relation.RelationByName(relationName)
		join := joinType(relation.JoinKey, relation.Name, ar.Relation.Name, relation.Columns, relation.TableName, relation.ForeignKey)
		if err != nil {
			return "", types.Relation{}, nil, fmt.Errorf("Relation doesnt exist")
		}

		ar.Join.OnJoin(join)

		return ar.Join.JoinKey(), relation, join, nil
	}

	return "", types.Relation{}, nil, fmt.Errorf("Relation doesnt exist")
}

func (query *SelectQuery[T]) Relation(relationName string, joinName string) *SelectQuery[T] {

	originKey, relation, join, err := query.relationByName(relationName, joinName)

	if err != nil {
		println(err.Error())
	}

	activeRelation := ActiveRelation{
		Relation: relation,
		Join:     join,
	}

	query.activeRelations = append(query.activeRelations, activeRelation)

	query.queryString += activeRelation.Join.JoinExpr(originKey, relation)

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

	queryString += "ORDER BY jk0__join_row"

	for _, relation := range query.activeRelations {
		queryString += fmt.Sprintf(", %s__join_row", relation.Join.JoinKey())
	}

	rows, err := db().Query(queryString)

	if err != nil {
		println(err.Error())
	}

	joins := []joins.Join{}

	for _, ar := range query.activeRelations {
		joins = append(joins, ar.Join)
	}

	sqlParser := sqlparser.NewSqlParser(query.Rows, query.origin, *rows)

	sqlParser.Parse(&query.Rows)

	return query.Rows
}