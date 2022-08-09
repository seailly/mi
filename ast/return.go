package ast

import (
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
