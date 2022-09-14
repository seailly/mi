package ast

import (
	"bytes"
	"fmt"

	"github.com/seailly/mi/token"
)

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("(%s%s)", pe.Operator, pe.Right.String()))

	return out.String()
}
