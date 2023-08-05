package driver

import (
	"encoding/json"
	"reflect"
	"strings"
)

var typeMap map[string]string = map[string]string{
	"string": "VARCHAR",
	"int":    "INT4",
}

type MockColumnType struct {
	databasetypename string
	name             string
}

func (mct *MockColumnType) DatabaseTypeName() string {
	return mct.databasetypename
}

func (mct *MockColumnType) Name() string {
	return mct.name
}

type MockRowsResult struct {
	rows         []interface{}
	currentIndex int
}

func (mrr *MockRowsResult) ColumnTypes() ([]ColumnType, error) {
	cts := []ColumnType{}

	rft := reflect.TypeOf(mrr.rows[0])

	for i := 0; i < rft.NumField(); i++ {
		cts = append(cts, &MockColumnType{
			databasetypename: typeMap[rft.Field(i).Type.Name()],
			name:             strings.ToLower(rft.Field(i).Name),
		})
	}

	return cts, nil
}

func (mrr *MockRowsResult) Next() bool {
	return mrr.currentIndex < len(mrr.rows)
}

func (mrr *MockRowsResult) Scan(dest ...any) error {
	byt, err := json.Marshal(&mrr.rows[mrr.currentIndex])

	if err != nil {
		return err
	}

	err = json.Unmarshal(byt, &dest)

	if err != nil {
		return err
	}

	return nil
}

type MockCommandResult struct {
}

type Mock struct {
	mockData        interface{}
	ExecutedQueries []string
}

func (m *Mock) Query(query string, args ...any) (RowsResult, error) {
	m.ExecutedQueries = append(m.ExecutedQueries, query)

	rft := reflect.TypeOf(m.mockData)

	rows := []interface{}{}

	if rft.Kind() == reflect.Slice {
		println("yes")

		rfv := reflect.ValueOf(m.mockData)

		for i := 0; i < rfv.Len(); i++ {
			rows = append(rows, rfv.Index(i).Interface())
		}
	}

	return &MockRowsResult{
		rows:         rows,
		currentIndex: 0,
	}, nil
}

func (m *Mock) Exec(query string, args ...any) (CommandResult, error) {
	m.ExecutedQueries = append(m.ExecutedQueries, query)

	return &MockCommandResult{}, nil
}

func NewMock(mockData interface{}) *Mock {
	return &Mock{
		mockData: mockData,
	}
}
