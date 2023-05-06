package types

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
}
