package ast

import (
	"github.com/seailly/mi/token"
)

// MutStatement
type MutStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// statementNode
func (ms *MutStatement) statementNode() {}

// TokenLiteral Return literal from node token
func (ms *MutStatement) TokenLiteral() string {
	return ms.Token.Literal
}
