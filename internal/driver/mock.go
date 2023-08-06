package driver

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

	cts = append(cts, &MockColumnType{
		databasetypename: typeMap["int"],
		name:             "jk0__join_row",
	})

	rft := reflect.TypeOf(mrr.rows[0])

	for i := 0; i < rft.NumField(); i++ {
		if rft.Field(i).Name != "baseModel" {
			cts = append(cts, &MockColumnType{
				databasetypename: typeMap[rft.Field(i).Type.Name()],
				name:             fmt.Sprintf("jk0__%s", strings.ToLower(rft.Field(i).Name)),
			})
		}
	}

	return cts, nil
}

func (mrr *MockRowsResult) Next() bool {
	return mrr.currentIndex < len(mrr.rows)
}

func (mrr *MockRowsResult) Scan(dest ...any) error {
	currRow := mrr.rows[mrr.currentIndex]

	rft := reflect.TypeOf(currRow)
	rfv := reflect.ValueOf(currRow)

	destIndex := 0

	byt, err := json.Marshal(sql.NullInt64{
		Int64: int64(mrr.currentIndex + 1),
		Valid: true,
	})

	if err != nil {
		return err
	}

	if destIndex < len(dest) {
		err = json.Unmarshal(byt, &dest[destIndex])
	}

	if err != nil {
		return err
	}

	destIndex = 1

	for i := 0; i < rft.NumField(); i++ {
		if rft.Field(i).IsExported() {

			var byt []byte
			var err error

			switch rfv.Field(i).Kind() {
			case reflect.String:
				byt, err = json.Marshal(sql.NullString{
					String: rfv.Field(i).String(),
					Valid:  true,
				})
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
				byt, err = json.Marshal(sql.NullInt64{
					Int64: rfv.Field(i).Int(),
					Valid: true,
				})
			}

			if err != nil {
				return err
			}

			if destIndex < len(dest) {
				err = json.Unmarshal(byt, &dest[destIndex])
			}

			if err != nil {
				return err
			}

			destIndex += 1
		}
	}

	mrr.currentIndex += 1

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
