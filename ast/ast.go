package ast

import (
	"bytes"
)

// Node A node in the AST
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement Refering to code the returns a statement
type Statement interface {
	Node
	statementNode()
}

// Expression Refering to code the returns a expression
type Expression interface {
	Node
	expressionNode()
}
