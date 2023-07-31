package gormy

import (
	"fmt"
)

type Relation struct {
	Name       string
	Type       string
	How        string
	Key        string
	ForeignKey string
	TableName  string
	JoinKey    string
	Columns    []Column
	Model      interface{}
	TagData    map[string]string
	Relations  []Relation
}

func (r *Relation) RelationByName(relationName string) (Relation, error) {
	for _, relation := range r.Relations {
		if relation.Name == relationName {
			return relation, nil
		}
	}

	return Relation{}, fmt.Errorf("Relation doesnt exist")
}
