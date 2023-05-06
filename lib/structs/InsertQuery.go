package structs

type InsertQuery[T any] struct {
	Query[T]
}

func (query *InsertQuery[T]) Exec() string {
	return query.queryString
}
