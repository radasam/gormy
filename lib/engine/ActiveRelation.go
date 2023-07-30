package engine

import (
	"github.com/radasam/gormy/lib/joins"
	"github.com/radasam/gormy/lib/types"
)

type ActiveRelation struct {
	Relation types.Relation
	Join     joins.Join
	JoinKey  string
}
