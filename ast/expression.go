package ast

import (
	"github.com/seailly/mi/token"
)

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// statementNode
func (es *ExpressionStatement) statementNode() {}

// TokenLiteral
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
