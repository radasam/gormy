package engine

import (
	"fmt"
	"gormy/lib/fileutils"
)

type importedColumn struct {
	Name      string `json:"sql_name"`
	ModelName string `json:"model_name"`
	SqlType   string `json:"sql_type"`
	ModelType string `json:"model_type"`
}

type importedTable struct {
	tableName  string
	structName string
	columns    []importedColumn
}

func (t *importedTable) ToFile(path string) error {
	w, err := fileutils.NewWriter(fmt.Sprintf("%s/%s.go", path, t.structName), t.structName)

	if err != nil {
		return fmt.Errorf("writing to file: %w", err)
	}

	for _, ic := range t.columns {
		w.Append(ic.ModelName, ic.ModelType, ic.SqlType, ic.Name)
	}

	err = w.Close()

	if err != nil {
		return fmt.Errorf("writing to file: %w", err)
	}

	return nil
}
