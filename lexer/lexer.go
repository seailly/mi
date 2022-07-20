package lexer

import (
	"github.com/seailly/mi/token"
)

const ASCIINul = 0

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar Set the next character and advance position in the input string
func (l *Lexer) readChar() {
	// Check if we have reached the end of the input
	if l.readPosition >= len(l.input) {
		l.ch = ASCIINul // ASCII code for "NUL"
	} else {
		l.ch = l.input[l.readPosition] // Advance to next character
	}

	l.position = l.readPosition
	l.readPosition += 1 // Advance next read position
}

// NextToken Match ASCII with token.TokenType
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ASCIINul:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar() // Move read position
	return tok
}

// newToken Create a token.Token based on tokenType and character
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
