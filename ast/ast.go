package ast

import (
	"github.com/seailly/mi/token"
)

// Node A node in the AST
type Node interface {
	TokenLiteral() string
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

// Program This is the root Node for every AST
type Program struct {
	Statements []Statement
}

// TokenLiteral
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

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
