package types

type Column struct {
	StructName   string
	Name         string
	DataType     string
	IsRelation   bool
	RelationName string
	JoinKey      string
}
