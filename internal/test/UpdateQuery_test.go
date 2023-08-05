package test

import (
	"testing"

	"github.com/radasam/gormy"
	"github.com/radasam/gormy/internal/driver"
	"github.com/radasam/gormy/internal/test/testmodels"
)

func TestSimpleUpdate(t *testing.T) {
	testData := []testmodels.SimpleModel{
		{Name: "Steve", Age: 54}, {Name: "Mary", Age: 32},
	}

	mock := driver.NewMock(testData)

	gormy.UseMockClient(mock)

	_, err := gormy.Model(testmodels.SimpleModel{}).Query().Update().Set("? = '?'", "name", "Dan").Where("? = '?'", "name", "Steve").Exec()

	if err != nil {
		t.Errorf("unexpected update error: %s", err.Error())
	}

	expQuery := "UPDATE simplemodel as jk0\r\nSET name = 'Dan'\r\nWHERE name = 'Steve'\r\n"

	if expQuery != mock.ExecutedQueries[0] {
		t.Errorf("expected query : %s got %s", expQuery, mock.ExecutedQueries[0])
	}
}

func TestMultipleSetUpdate(t *testing.T) {
	testData := []testmodels.SimpleModel{
		{Name: "Steve", Age: 54}, {Name: "Mary", Age: 32},
	}

	mock := driver.NewMock(testData)

	gormy.UseMockClient(mock)

	_, err := gormy.Model(testmodels.SimpleModel{}).Query().Update().
		Set("? = '?'", "name", "Dan").
		Set("? = ?", "age", "52").
		Where("? = '?'", "name", "Steve").Exec()

	if err != nil {
		t.Errorf("unexpected update error: %s", err.Error())
	}

	expQuery := "UPDATE simplemodel as jk0\r\nSET name = 'Dan'\r\n, age = 52\r\nWHERE name = 'Steve'\r\n"

	if expQuery != mock.ExecutedQueries[0] {
		t.Errorf("expected query : %s got %s", expQuery, mock.ExecutedQueries[0])
	}
}

func TestMultipleWhereUpdate(t *testing.T) {
	testData := []testmodels.SimpleModel{
		{Name: "Steve", Age: 54}, {Name: "Mary", Age: 32},
	}

	mock := driver.NewMock(testData)

	gormy.UseMockClient(mock)

	_, err := gormy.Model(testmodels.SimpleModel{}).Query().Update().
		Set("? = '?'", "name", "Dan").
		Set("? = ?", "age", "52").
		Where("? = '?'", "name", "Steve").
		Where("? = ?", "age", "54").
		Exec()

	if err != nil {
		t.Errorf("unexpected update error: %s", err.Error())
	}

	expQuery := "UPDATE simplemodel as jk0\r\nSET name = 'Dan'\r\n, age = 52\r\nWHERE name = 'Steve'\r\nAND age = 54\r\n"

	if expQuery != mock.ExecutedQueries[0] {
		t.Errorf("expected query : %s got %s", expQuery, mock.ExecutedQueries[0])
	}
}
