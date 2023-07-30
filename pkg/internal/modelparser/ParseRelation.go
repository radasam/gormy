package modelparser

import (
	"fmt"
	"github.com/radasam/gormy/pkg/internal/types"
	"reflect"
	"strings"
)

func ParseRelation(relationTag string, structName string, relationModel reflect.Type, relationPrefix string) (types.Relation, error) {
	relationCount := 0
	relation := types.Relation{}
	relationColumns := []types.Column{}
	nestedRelations := []types.Relation{}
	tagData := map[string]string{}

	tagItems := strings.Split(strings.Split(strings.Split(relationTag, "gormy:\"")[1], "\"")[0], ",")

	for _, tagItem := range tagItems {
		name := strings.Split(tagItem, ":")[0]

		if name != "relation" {
			value := strings.Split(tagItem, ":")[1]

			switch name {
			case "how":
				relation.How = strings.ToLower(value)
			case "on":
				key := strings.Split(value, "=")[0]
				foreignKey := strings.Split(value, "=")[1]

				relation.Key = key
				relation.ForeignKey = foreignKey
			default:
				tagData[name] = value
				// return relation, fmt.Errorf("ParseRelationError: Unidentified option - %s", name)
			}
		} else {
			relation.Type = name
		}
	}

	if relationModel.Kind() == reflect.Slice {
		relationModel = relationModel.Elem()
	}

	relation.TagData = tagData

	for i := 0; i < relationModel.NumField(); i++ {
		name := relationModel.Field(i).Name
		hasRelation := strings.Contains(relationModel.Field(i).Tag.Get("gormy"), "relation:")
		if name != "baseModel" && !hasRelation {
			column, _, err := ParseColumn(string(relationModel.Field(i).Tag), relationModel.Field(i).Name, relationModel.Field(i).Type, relationPrefix, relationCount)

			if err != nil {
				println(err.Error())
			}

			relationColumns = append(relationColumns, *column)
		} else if !hasRelation {
			tableName, err := ParseConfig(string(relationModel.Field(i).Tag))

			if err != nil {
				println(err.Error())
			}

			relation.TableName = tableName
		} else {
			relationCount += 1
			nestedRelation, err := ParseRelation(string(relationModel.Field(i).Tag), relationModel.Field(i).Name, relationModel.Field(i).Type, fmt.Sprintf("%s_jk%d", relationPrefix, relationCount))
			nestedRelation.JoinKey = fmt.Sprintf("%s_jk%d", relationPrefix, relationCount)

			if err != nil {
				return types.Relation{}, err
			}

			nestedRelations = append(nestedRelations, nestedRelation)
		}
	}

	relation.Name = structName
	relation.Columns = relationColumns
	relation.Relations = nestedRelations

	return relation, nil
}
