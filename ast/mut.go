package ast

import (
	"bytes"

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

// String
func (ms *MutStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ms.TokenLiteral() + " ")
	out.WriteString(ms.Name.String())
	out.WriteString(" = ")

	if ms.Value != nil {
		out.WriteString(ms.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
