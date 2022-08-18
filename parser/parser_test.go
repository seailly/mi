package parser

import (
	"fmt"
	"testing"

	"github.com/seailly/mi/ast"
	"github.com/seailly/mi/lexer"
	"github.com/stretchr/testify/require"
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
	require.NotNil(t, p)

	checkParserErrors(t, p)
	require.Len(t, program.Statements, 3)

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
	require.NotNil(t, program)
	require.Len(t, p.errors, 3)
}

func TestReturnStatement(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	require.NotNil(t, program)

	// TODO: Test error case
	require.Equal(t, len(program.Statements), 3)

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got %T", stmt)
			continue
		}

		require.Equal(t, returnStmt.TokenLiteral(), "return")
	}
}

func testMutStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "mut" {
		t.Errorf("s.TokenLiteral not 'mut'. got %q", s.TokenLiteral())
		return false
	}

	mutStmt, ok := s.(*ast.MutStatement)
	require.True(t, ok, "s not *ast.MutStatement. got %T", s)

	require.Equalf(t, mutStmt.Name.Value, name, "muStmt.Name.Value not '%s'. got %s", name, mutStmt.Name.Value)
	require.Equalf(t, mutStmt.Name.TokenLiteral(), name, "s.Name not '%s'. got %s", name, mutStmt.Name)

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

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.Len(t, program.Statements, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	require.True(t, ok, fmt.Sprintf("program.Statements[0] is not ast.ExpressionStatement. got %T",
		program.Statements[0]))

	ident, ok := stmt.Expression.(*ast.Identifier)
	require.True(t, ok)
	require.Equal(t, ident.Value, "foobar")
	require.Equal(t, ident.TokenLiteral(), "foobar")
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.Len(t, program.Statements, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	require.True(t, ok)

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	require.True(t, ok)
	require.Equal(t, literal.Value, int64(5))
	require.Equal(t, literal.TokenLiteral(), "5")
}
