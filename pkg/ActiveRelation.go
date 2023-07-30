package gormy

import (
	"github.com/radasam/gormy/pkg/internal/joins"
	"github.com/radasam/gormy/pkg/internal/types"
)

type ActiveRelation struct {
	Relation types.Relation
	Join     joins.Join
	JoinKey  string
}
