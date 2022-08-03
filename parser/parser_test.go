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
