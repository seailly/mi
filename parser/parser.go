package parser

import (
	"fmt"
	"github.com/seailly/mi/ast"
	"github.com/seailly/mi/lexer"
	"github.com/seailly/mi/token"
)

// Parser
type Parser struct {
	l *lexer.Lexer

	errors []string

	curToken  token.Token
	peekToken token.Token
}

// New Returns a Parser with setup lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	// Reads two tokens so both curToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

// Errors Return list of errors
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)

	p.errors = append(p.errors, msg)
}

// nextToken
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// curTokenIs Checks if the current token matches the conditional token
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs Seeks forward 1 to see if the next token matches the conditional token
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek Asserts if the next token matches the conditional token
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// ParseProgram Creates the root AST node interating until EOF is found
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// parseStatement
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.MUT:
		return p.parseMutStatement()
	default:
		return nil
	}
}

// parseMutStatement
func (p *Parser) parseMutStatement() *ast.MutStatement {
	stmt := &ast.MutStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: Parse expression

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
