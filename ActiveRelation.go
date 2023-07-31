package gormy

type ActiveRelation struct {
	Relation Relation
	Join     Join
	JoinKey  string
}
