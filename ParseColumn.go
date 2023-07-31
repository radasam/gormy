package gormy

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func ParseColumn(columnTag string, structName string, goType reflect.Type, relationPrefix string, relationCount int) (*Column, *Relation, error) {

	IsRelation := false
	column := Column{}
	relation := Relation{}
	var err error

	if columnTag == "" {
		return nil, nil, errors.New("No data type provided")
	}

	tagItems := strings.Split(strings.Split(strings.Split(columnTag, "gormy:\"")[1], "\"")[0], ",")

	for _, tagItem := range tagItems {
		if strings.Contains(tagItem, ":") && !IsRelation {
			name := strings.Split(tagItem, ":")[0]
			value := strings.Split(tagItem, ":")[1]

			switch name {
			case "name":
				column.Name = strings.ToLower(value)
			case "relation":
				IsRelation = true

				relation, err = ParseRelation(columnTag, structName, goType, fmt.Sprintf("%s_jk%d", relationPrefix, relationCount+1))

				relation.JoinKey = fmt.Sprintf("%s_jk%d", relationPrefix, relationCount+1)

				if err != nil {
					return nil, nil, fmt.Errorf("parsing relation: %w", err)
				}
			default:
				return &column, nil, fmt.Errorf("Unidentified option - %s", name)
			}
		} else {
			column.DataType = tagItem
		}
	}

	if IsRelation {
		return nil, &relation, nil
	} else {
		if column.DataType == "" {
			return nil, nil, errors.New("No data type provided")
		}

		column.StructName = structName

		if column.Name == "" {
			column.Name = structName
		}

		return &column, nil, nil
	}

}
