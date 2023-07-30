package gormy

import "strings"

func DefaultSchemaMap() map[string]string {

	dmap := map[string]string{
		"varchar": "string",
		"int4":    "int",
	}

	return dmap
}

func DefaultNameMap(columnName string) string {
	comp := strings.Split(columnName, "_")

	for i, v := range comp {
		comp[i] = strings.Title(v)
	}

	return strings.Join(comp, "")
}
