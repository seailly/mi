package ast

import (
	"github.com/seailly/mi/token"
)

// Identifier
type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

// expressionNode
func (i *Identifier) expressionNode() {}

// TokenLiteral Return literal from node token
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String
func (i *Identifier) String() string {
	return i.Value
}
