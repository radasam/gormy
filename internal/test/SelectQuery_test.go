package test

import (
	"encoding/json"
	"testing"

	"github.com/radasam/gormy"
	"github.com/radasam/gormy/internal/driver"
	"github.com/radasam/gormy/internal/test/testmodels"
)

func TestSimpleSelect(t *testing.T) {
	testData := []testmodels.SimpleModel{
		{Name: "Steve", Age: 54}, {Name: "Mary", Age: 32},
	}

	mock := driver.NewMock(testData)

	gormy.UseMockClient(mock)

	sm, err := gormy.Model(testmodels.SimpleModel{}).Query().Select().Exec()

	if err != nil {
		t.Errorf("unexpected update error: %s", err.Error())
	}

	expQuery := "SELECT jk0.Name as jk0__Name,jk0.Age as jk0__Age \r\nFROM simplemodel as jk0\r\n"

	if expQuery != mock.ExecutedQueries[0] {
		t.Errorf("expected query : %s got %s", expQuery, mock.ExecutedQueries[0])
	}

	ebyt, err := json.Marshal(testData)

	if err != nil {
		t.Errorf("unexpected update error: %s", err.Error())
	}

	abyt, err := json.Marshal(sm)

	if err != nil {
		t.Errorf("unexpected update error: %s", err.Error())
	}

	if string(ebyt) != string(abyt) {
		t.Errorf("expected result %s got %s", string(ebyt), string(abyt))
	}

}

func TestSelectWhere(t *testing.T) {
	testData := []testmodels.SimpleModel{
		{Name: "Steve", Age: 54}, {Name: "Mary", Age: 32},
	}

	mock := driver.NewMock(testData)

	gormy.UseMockClient(mock)

	sm, err := gormy.Model(testmodels.SimpleModel{}).Query().Select().Where("? = '?'", "name", "Steve").Exec()

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	expQuery := "SELECT jk0.Name as jk0__Name,jk0.Age as jk0__Age \r\nFROM simplemodel as jk0\r\nWHERE name = 'Steve'\r\n"

	if expQuery != mock.ExecutedQueries[0] {
		t.Errorf("expected query : %s got %s", expQuery, mock.ExecutedQueries[0])
	}

	ebyt, err := json.Marshal(testData)

	if err != nil {
		t.Errorf("unexpected update error: %s", err.Error())
	}

	abyt, err := json.Marshal(sm)

	if err != nil {
		t.Errorf("unexpected update error: %s", err.Error())
	}

	if string(ebyt) != string(abyt) {
		t.Errorf("expected result %s got %s", string(ebyt), string(abyt))
	}

}

func TestSelectMultipleWhere(t *testing.T) {
	testData := []testmodels.SimpleModel{
		{Name: "Steve", Age: 54}, {Name: "Mary", Age: 32},
	}

	mock := driver.NewMock(testData)

	gormy.UseMockClient(mock)

	sm, err := gormy.Model(testmodels.SimpleModel{}).Query().Select().Where("? = '?'", "name", "Steve").Where("? = ?", "age", "54").Exec()

	if err != nil {
		t.Errorf("unexpected update error: %s", err.Error())
	}

	expQuery := "SELECT jk0.Name as jk0__Name,jk0.Age as jk0__Age \r\nFROM simplemodel as jk0\r\nWHERE name = 'Steve'\r\nAND age = 54\r\n"

	if expQuery != mock.ExecutedQueries[0] {
		t.Errorf("expected query : %s got %s", expQuery, mock.ExecutedQueries[0])
	}

	ebyt, err := json.Marshal(testData)

	if err != nil {
		t.Errorf("unexpected update error: %s", err.Error())
	}

	abyt, err := json.Marshal(sm)

	if err != nil {
		t.Errorf("unexpected update error: %s", err.Error())
	}

	if string(ebyt) != string(abyt) {
		t.Errorf("expected result %s got %s", string(ebyt), string(abyt))
	}

}
