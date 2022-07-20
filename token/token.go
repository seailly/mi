package token

// TokenType
type TokenType string

// Token
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals

	// Ident
	IDENT = "IDENT" // add, foobar, x, y, ...
	// Int
	INT = "INT" // 1343456”

	// Operators

	// Assign
	ASSIGN = "="

	// Plus
	PLUS = "+"

	// Delimiters

	// Comma
	COMMA = ","

	// Semicolon
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords

	// Function
	FUNCTION = "FUNCTION"

	// Mut
	MUT = "MUT"
)
