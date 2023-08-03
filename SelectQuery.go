package gormy

import (
	"fmt"
	"strings"
)

type SelectQuery[T any] struct {
	Query[T]
	selected        []Column
	activeRelations []ActiveRelation
	errored         error
}

func (query *SelectQuery[T]) Column(columnName string) *SelectQuery[T] {

	column, err := query.columnByName(columnName)

	if err != nil {
		query.errored = fmt.Errorf("selectec column doesnt exist: %s", columnName)
		return query
	}

	query.selected = append(query.selected, column)

	return query
}

func (query *SelectQuery[T]) Where(expr string, columnName string, value string) *SelectQuery[T] {

	cleanExpr := strings.ReplaceAll(expr, "?", "%s")

	query.queryString = query.queryString + fmt.Sprintf("WHERE "+cleanExpr+"\r\n", columnName, value)

	return query
}

func (query *SelectQuery[T]) relationByName(relationName string, joinName string) (string, Relation, Join, error) {

	joinType, err := gc.RegisteredJoins.ByName(joinName)

	if err != nil {
		return "", Relation{}, nil, fmt.Errorf("Relation doesnt exist")
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
			return "", Relation{}, nil, fmt.Errorf("Relation doesnt exist")
		} else {
			ar.Join.OnJoin(join)
			return ar.Join.JoinKey(), relation, join, nil
		}

	}

	return "", Relation{}, nil, fmt.Errorf("Relation doesnt exist")
}

func (query *SelectQuery[T]) Relation(relationName string, joinName string) *SelectQuery[T] {

	originKey, relation, join, err := query.relationByName(relationName, joinName)

	if err != nil {
		query.errored = err
		return query
	}

	activeRelation := ActiveRelation{
		Relation: relation,
		Join:     join,
	}

	query.activeRelations = append(query.activeRelations, activeRelation)

	query.queryString += activeRelation.Join.JoinExpr(originKey, relation)

	return query
}

func (query *SelectQuery[T]) Exec() ([]T, error) {

	if query.errored != nil {
		return []T{}, query.errored
	}

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

	if len(query.activeRelations) > 0 {
		queryString += "ORDER BY jk0__join_row"
	}

	for _, relation := range query.activeRelations {
		if relation.Join.HasJoin() {
			queryString += fmt.Sprintf(", %s__join_row", relation.Join.JoinKey())
		}
	}

	rows, err := gc.conn.Query(queryString)

	if err != nil {
		return []T{}, fmt.Errorf("executing query: %w", err)
	}

	sqlParser := newSqlParser(query.Rows, query.origin, *rows)

	sqlParser.Parse(&query.Rows)

	return query.Rows, nil
}
