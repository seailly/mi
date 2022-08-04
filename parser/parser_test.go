package parser

import (
	"testing"

	"github.com/seailly/mi/ast"
	"github.com/seailly/mi/lexer"
)

func TestMutStatement(t *testing.T) {
	input := `
mut x = 5;
mut y = 10;
mut foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifer string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testMutStatement(t, stmt, tt.expectedIdentifer) {
			return
		}
	}
}

func TestMutStatement_CauseError(t *testing.T) {
	input := `
mut x 5;
mut = 10;
mut 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(p.errors) != 3 {
		t.Errorf("Error not caught by parser")
	}
}

func testMutStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "mut" {
		t.Errorf("s.TokenLiteral not 'mut'. got %q", s.TokenLiteral())
		return false
	}

	mutStmt, ok := s.(*ast.MutStatement)
	if !ok {
		t.Errorf("s not *ast.MutStatement. got %T", s)
		return false
	}

	if mutStmt.Name.Value != name {
		t.Errorf("muStmt.Name.Value not '%s'. got %s", name, mutStmt.Name.Value)
		return false
	}

	if mutStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got %s", name, mutStmt.Name)
		return false
	}

	return true
}

// checkParserErrors Fails test suite if errors are found in parser
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.errors

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}
