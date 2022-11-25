package object

type ObjectType string

const (
	INTEGER_OBJECT = "INTEGER"
	BOOLEAN_OBJECT = "BOOLEAN"
	NULL_OBJECT = "NULL"
	RETURN_VALUE_OBJECT = "RETURN_VALUE"
)

// Object Each value represents itself
type Object interface {
	Type() ObjectType
	Inspect() string
}