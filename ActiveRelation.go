package gormy

import (
	"github.com/radasam/gormy/internal/joins"
	"github.com/radasam/gormy/internal/types"
)

type ActiveRelation struct {
	Relation types.Relation
	Join     joins.Join
	JoinKey  string
}
