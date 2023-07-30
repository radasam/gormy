package engine

import (
	"gormy/lib/joins"
	"gormy/lib/types"
)

type ActiveRelation struct {
	Relation types.Relation
	Join     joins.Join
	JoinKey  string
}
