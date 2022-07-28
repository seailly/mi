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

	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

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
	FUNCTION = "FN"

	// Mut
	MUT = "MUT"

	FALSE  = "FALSE"
	TRUE   = "TRUE"
	RETURN = "RETURN"
	IF     = "IF"
	ELSE   = "ELSE"
)

// keywords
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"mut":    MUT,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
}

// LookupIdent Find keyword TokenType by string
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
