package utils

import (
	"database/sql"
	"encoding/json"
	"strings"
)

func RowsToJson(rows sql.Rows, joinKeys []string, joinNames []string, joinTypes map[string]string) (string, error) {

	columnTypes, err := rows.ColumnTypes()

	if err != nil {
		return "", err
	}

	count := len(columnTypes)
	rowMap := map[int]map[string]interface{}{}
	repeatRowMap := map[int]int{}
	joinKeyArrayMap := map[int]map[string][]map[string]interface{}{}
	rowNumber := 0

	for rows.Next() {

		scanArgs := make([]interface{}, count)

		for i, v := range columnTypes {
			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
				break
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
				break
			case "INT4", "INT8":
				scanArgs[i] = new(sql.NullInt64)
				break
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		err := rows.Scan(scanArgs...)

		if err != nil {
			return "", err
		}

		joinKeyMap := map[string]map[string]interface{}{}

		if columnTypes[0].Name() == "gor_join_row" {
			if z, ok := (scanArgs[0]).(*sql.NullInt64); ok {
				rowNumber = int(z.Int64)
			}

			if _, ok := joinKeyArrayMap[rowNumber]; !ok {
				joinKeyArrayMap[rowNumber] = map[string][]map[string]interface{}{}
			}

			if _, ok := repeatRowMap[rowNumber]; !ok {
				repeatRowMap[rowNumber] = 0
			} else {
				repeatRowMap[rowNumber] += 1
			}

		} else {
			rowNumber += 1
		}

		if _, ok := rowMap[rowNumber]; !ok {
			rowMap[rowNumber] = map[string]interface{}{}
		}

		for _, v := range joinKeys {
			if joinTypes[v] == "onetoone" {
				joinKeyMap[v] = map[string]interface{}{}
			} else {
				if _, ok := joinKeyArrayMap[rowNumber][v]; !ok {
					joinKeyArrayMap[rowNumber][v] = []map[string]interface{}{}
				}

				joinKeyArrayMap[rowNumber][v] = append(joinKeyArrayMap[rowNumber][v], make(map[string]interface{}))
			}
		}

		for i, v := range columnTypes {

			if v.Name() != "gor_join_row" {
				key := strings.Split(v.Name(), "__")[0]
				name := strings.Split(v.Name(), "__")[1]

				if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
					if key == "jk0" {
						rowMap[rowNumber][name] = z.Bool
					} else {
						if joinTypes[key] == "onetoone" {
							joinKeyMap[key][name] = z.Bool
						} else {
							joinKeyArrayMap[rowNumber][key][repeatRowMap[rowNumber]][name] = z.Bool
						}
					}
					continue
				}

				if z, ok := (scanArgs[i]).(*sql.NullString); ok {
					if key == "jk0" {
						rowMap[rowNumber][name] = z.String
					} else {
						if joinTypes[key] == "onetoone" {
							joinKeyMap[key][name] = z.String
						} else {
							joinKeyArrayMap[rowNumber][key][repeatRowMap[rowNumber]][name] = z.String
						}
					}
					continue
				}

				if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
					if key == "jk0" {
						rowMap[rowNumber][name] = z.Int64
					} else {
						if joinTypes[key] == "onetoone" {
							joinKeyMap[key][name] = z.Int64
						} else {
							joinKeyArrayMap[rowNumber][key][repeatRowMap[rowNumber]][name] = z.Int64
						}
					}
					continue
				}

				if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
					if key == "jk0" {
						rowMap[rowNumber][name] = z.Float64
					} else {
						if joinTypes[key] == "onetoone" {
							joinKeyMap[key][name] = z.Float64
						} else {
							joinKeyArrayMap[rowNumber][key][repeatRowMap[rowNumber]][name] = z.Float64
						}

					}
					continue
				}

				if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
					if key == "jk0" {
						rowMap[rowNumber][name] = z.Int32
					} else {
						if joinTypes[key] == "onetoone" {
							joinKeyMap[key][name] = z.Int32
						} else {
							joinKeyArrayMap[rowNumber][key][repeatRowMap[rowNumber]][name] = z.Int32
						}

					}
					continue
				}

				if key == "jk0" {
					rowMap[rowNumber][name] = scanArgs[i]
				} else {
					if joinTypes[key] == "onetoone" {
						joinKeyMap[key][name] = scanArgs[i]
					} else {
						joinKeyArrayMap[rowNumber][key][repeatRowMap[rowNumber]][name] = scanArgs[i]
					}
				}
			}

		}

		for i, v := range joinKeys {
			if joinTypes[v] == "onetoone" {
				rowMap[rowNumber][joinNames[i]] = joinKeyMap[v]
			} else {
				rowMap[rowNumber][joinNames[i]] = joinKeyArrayMap[rowNumber][v]
			}
		}
	}

	values := make([]map[string]interface{}, 0, len(rowMap))
	for k := range rowMap {
		values = append(values, rowMap[k])
	}

	z, err := json.Marshal(values)

	if err != nil {
		return "", err
	}

	return string(z), err

}
