package ast

import (
	"bytes"

	"github.com/seailly/mi/token"
)

// ReturnStatement
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// statementNode
func (rs *ReturnStatement) statementNode() {}

// TokenLiteral
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
