package modelparser

import (
	"errors"
	"fmt"
	"gormy/lib/types"
	"reflect"
	"strings"
)

func ParseColumn(columnTag string, structName string, goType reflect.Type, relationPrefix string, relationCount int) (*types.Column, *types.Relation, error) {

	IsRelation := false
	column := types.Column{}
	relation := types.Relation{}
	var err error

	if columnTag == "" {
		return nil, nil, errors.New("ParseColumnError: No data type provided.")
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
					return nil, nil, err
				}
			default:
				return &column, nil, fmt.Errorf("ParseColumnError: Unidentified option - %s", name)
			}
		} else {
			column.DataType = tagItem
		}
	}

	if IsRelation {
		return nil, &relation, nil
	} else {
		if column.DataType == "" {
			return nil, nil, errors.New("ParseColumnError: No data type provided.")
		}

		column.StructName = structName

		if column.Name == "" {
			column.Name = structName
		}

		return &column, nil, nil
	}

}
