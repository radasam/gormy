package driver

type ColumnType interface {
	DatabaseTypeName() string
	Name() string
}

type RowsResult interface {
	ColumnTypes() ([]ColumnType, error)
	Next() bool
	Scan(...any) error
}

type CommandResult interface {
}

type Driver interface {
	Query(string, ...any) (RowsResult, error)
	Exec(string, ...any) (CommandResult, error)
}
