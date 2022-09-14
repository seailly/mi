package ast

import (
	"bytes"
	"fmt"

	"github.com/seailly/mi/token"
)

type InfixExpression struct {
	Token    token.Token // Operator token eg: +
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode() {

}

func (oe *InfixExpression) TokenLiteral() string {
	return oe.Token.Literal
}

func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("(%s %s %s)", oe.Left.String(), oe.Operator, oe.Right.String()))
	return out.String()
}
