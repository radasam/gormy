package gormy

import (
	"fmt"
	"strings"
)

func ParseConfig(configTag string) (string, error) {
	tableName := ""

	tagItems := strings.Split(strings.Split(strings.Split(configTag, "gormy:\"")[1], "\"")[0], ",")

	for _, tagItem := range tagItems {
		if strings.Contains(tagItem, ":") {
			name := strings.Split(tagItem, ":")[0]

			switch name {
			default:
				return "", fmt.Errorf("Unidentified option - %s", name)
			}
		} else {
			tableName = tagItem
		}
	}

	return tableName, nil
}
