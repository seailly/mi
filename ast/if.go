package ast

import (
	"bytes"
	"fmt"

	"github.com/seailly/mi/token"
)

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {

}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("if %s %s", ie.Condition.String(), ie.Consequence.String()))
	if ie.Alternative != nil {
		out.WriteString(fmt.Sprintf("else %s", ie.Alternative.String()))
	}

	return out.String()
}
