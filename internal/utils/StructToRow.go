package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

func StructToRow(field reflect.Value) string {
	switch field.Kind() {
	case reflect.String:
		return fmt.Sprintf("'%s'", field.String())
	case reflect.Int64:
		return strconv.FormatInt(field.Int(), 10)
	case reflect.Int:
		return strconv.FormatInt(field.Int(), 10)
	default:
		return field.String()
	}
}
