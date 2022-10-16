package object

type ObjectType string

const (
	INTEGER_OBJECT = "INTEGER"
	BOOLEAN_OBJECT = "BOOLEAN"
	NULL_OBJECT = "NULL"
)

// Object Each value represents itself
type Object interface {
	Type() ObjectType
	Inspect() string
}