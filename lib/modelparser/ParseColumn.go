package modelparser

import (
	"errors"
	"gormy/lib/structs"
	"strings"
)

func ParseColumn(columnTag string, structName string) (structs.Column, error) {

	column := structs.Column{}

	if columnTag == "" {
		return column, errors.New("ParseColumnError: No data type provided.")
	}

	tagItems := strings.Split(strings.Split(strings.Split(columnTag, "gormy:\"")[1], "\"")[0], ",")

	for _, tagItem := range tagItems {
		if strings.Contains(tagItem, ":") {
			name := strings.Split(tagItem, ":")[0]
			value := strings.Split(tagItem, ":")[1]

			switch name {
			case "name":
				column.Name = value
			default:
				return column, errors.New("ParseColumnError: Unidentified option.")
			}
		} else {
			column.DataType = tagItem
		}
	}

	if column.DataType == "" {
		return column, errors.New("ParseColumnError: No data type provided.")
	}

	if column.Name == "" {
		column.Name = structName
	}

	return column, nil
}
