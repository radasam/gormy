package driver

import (
	"database/sql"
	"fmt"
)

type PostgresColumnType struct {
	columntype *sql.ColumnType
}

func (pct *PostgresColumnType) DatabaseTypeName() string {
	return pct.columntype.DatabaseTypeName()
}

func (pct *PostgresColumnType) Name() string {
	return pct.columntype.Name()
}

type PostgresRowsResult struct {
	rows *sql.Rows
}

func (prr *PostgresRowsResult) ColumnTypes() ([]ColumnType, error) {
	cts, err := prr.rows.ColumnTypes()

	if err != nil {
		return []ColumnType{}, err
	}

	pcts := []ColumnType{}

	for _, ct := range cts {
		pcts = append(pcts, &PostgresColumnType{
			columntype: ct,
		})
	}

	return pcts, nil
}

func (prr *PostgresRowsResult) Next() bool {
	return prr.rows.Next()
}

func (prr *PostgresRowsResult) Scan(dest ...any) error {
	return prr.rows.Scan(dest)
}

type Postgres struct {
	conn *sql.DB
}

func (p *Postgres) Query(query string, args ...any) (RowsResult, error) {
	rows, err := p.conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	return &PostgresRowsResult{
		rows,
	}, nil
}

func (p *Postgres) Exec(query string, args ...any) (CommandResult, error) {
	return p.conn.Exec(query, args)
}

func NewPostgres(connString string) (*Postgres, error) {
	conn, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	return &Postgres{
		conn: conn,
	}, nil
}
