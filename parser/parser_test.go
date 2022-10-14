package parser

import (
	"fmt"
	"testing"

	"github.com/seailly/mi/ast"
	"github.com/seailly/mi/lexer"
	"github.com/stretchr/testify/require"
)

// TestMutStatements
func TestMutStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"mut x = 5;", "x", 5},
		{"mut y = true;", "y", true},
		{"mut foobar = y;", "foobar", "y"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		require.Len(t, program.Statements, 1)

		stmt := program.Statements[0]
		testMutStatement(t, stmt, tt.expectedIdentifier)

		val := stmt.(*ast.MutStatement).Value
		testLiteralExpression(t, val, tt.expectedValue)
	}
}

func TestMutStatement_CauseError(t *testing.T) {
	input := `
mut x 5;
mut = 10;
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

func TestBooleanExpression(t *testing.T) {
	input := "true;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.Len(t, program.Statements, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	require.True(t, ok)

	literal, ok := stmt.Expression.(*ast.Boolean)
	require.True(t, ok)
	require.Equal(t, literal.Value, true)
	require.Equal(t, literal.TokenLiteral(), "true")
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue interface{}
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
		{"!true;", "!", true},
		{"!false;", "!", false},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		require.NotNil(t, program)
		checkParserErrors(t, p)
		require.Len(t, program.Statements, 1)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		require.True(t, ok)

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		require.True(t, ok)
		require.Equal(t, exp.Operator, tt.operator)
		testLiteralExpression(t, exp.Right, tt.integerValue)
	}
}

// TestParsingInfixExpressions
func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		require.NotNil(t, program)

		checkParserErrors(t, p)

		require.Len(t, program.Statements, 1)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		require.True(t, ok)

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		require.True(t, ok)

		testLiteralExpression(t, exp.Left, tt.leftValue)
		require.Equal(t, exp.Operator, tt.operator)
		testLiteralExpression(t, exp.Right, tt.rightValue)
	}
}

// testInfixExpression
func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) {
	opExp, ok := exp.(*ast.InfixExpression)

	require.True(t, ok)
	testLiteralExpression(t, opExp.Left, left)
	require.Equal(t, operator, opExp.Operator)
	testLiteralExpression(t, opExp.Right, right)
}

// TestOperatorPrecedenceParsing
func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()
		require.Equal(t, tt.expected, actual)
	}
}

// testMutStatement
func testMutStatement(t *testing.T, s ast.Statement, name string) {
	require.Equal(t, s.TokenLiteral(), "mut")

	mutStmt, ok := s.(*ast.MutStatement)
	require.True(t, ok, "s not *ast.MutStatement. got %T", s)

	require.Equalf(t, mutStmt.Name.Value, name, "muStmt.Name.Value not '%s'. got %s", name, mutStmt.Name.Value)
	require.Equalf(t, mutStmt.Name.TokenLiteral(), name, "s.Name not '%s'. got %s", name, mutStmt.Name)
}

// testLiteralExpression
func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) {
	switch v := expected.(type) {
	case int:
		testIntegerLiteral(t, exp, int64(v))
	case int64:
		testIntegerLiteral(t, exp, v)
	case bool:
		testBooleanLiteral(t, exp, v)
	case string:
		testIdentifier(t, exp, v)
	}
}

// testIntegerLiteral
func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) {
	integ, ok := il.(*ast.IntegerLiteral)
	require.True(t, ok)
	require.Equal(t, integ.Value, value)
	require.Equal(t, integ.TokenLiteral(), fmt.Sprintf("%d", value))
}

// testIdentifier
func testIdentifier(t *testing.T, exp ast.Expression, value string) {
	ident, ok := exp.(*ast.Identifier)
	require.True(t, ok)
	require.Equal(t, value, ident.Value)
	require.Equal(t, value, ident.TokenLiteral())
}

// testBooleanLiteral
func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) {
	bo, ok := exp.(*ast.Boolean)
	require.True(t, ok)
	require.Equal(t, value, bo.Value)
	require.Equal(t, fmt.Sprintf("%t", value), bo.TokenLiteral())
}

// TestIfExpression
func TestIfExpression(t *testing.T) {
	input := `if (x < y) { x }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.Len(t, program.Statements, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	require.True(t, ok)

	exp, ok := stmt.Expression.(*ast.IfExpression)
	require.True(t, ok)

	testInfixExpression(t, exp.Condition, "x", "<", "y")
	require.Len(t, exp.Consequence.Statements, 1)

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	require.True(t, ok)

	testIdentifier(t, consequence.Expression, "x")
	require.Nil(t, exp.Alternative)
}

// TestIfElseExpression
func TestIfElseExpression(t *testing.T) {
	input := `if (x < y) { x } else { y }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.Len(t, program.Statements, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	require.True(t, ok)

	exp, ok := stmt.Expression.(*ast.IfExpression)
	require.True(t, ok)

	testInfixExpression(t, exp.Condition, "x", "<", "y")
	require.Len(t, exp.Consequence.Statements, 1)

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	require.True(t, ok)

	testIdentifier(t, consequence.Expression, "x")
	require.Len(t, exp.Alternative.Statements, 1)

	alternative, ok := exp.Alternative.Statements[0].(*ast.ExpressionStatement)
	require.True(t, ok)

	testIdentifier(t, alternative.Expression, "y")
}

// TestFunctionLiteralParsing
func TestFunctionLiteralParsing(t *testing.T) {
	input := `fn(x, y) { x + y; }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.Len(t, program.Statements, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	require.True(t, ok)

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	require.True(t, ok)
	require.Len(t, function.Parameters, 2)

	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	require.Len(t, function.Body.Statements, 1)

	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)

	testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")
}

// TestFunctionParameterParsing catch edges cases from parsing function parameters
func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: "fn() {};", expectedParams: []string{}},
		{input: "fn(x) {};", expectedParams: []string{"x"}},
		{input: "fn(x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		function := stmt.Expression.(*ast.FunctionLiteral)
		require.Equal(t, len(function.Parameters), len(tt.expectedParams))

		for i, ident := range tt.expectedParams {
			testLiteralExpression(t, function.Parameters[i], ident)
		}
	}
}

// TestCallExpressionParsing
func TestCallExpressionParsing(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5)"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.Len(t, program.Statements, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	require.True(t, ok)

	exp, ok := stmt.Expression.(*ast.CallExpression)
	require.True(t, ok)

	testIdentifier(t, exp.Function, "add")
	require.Len(t, exp.Arguments, 3)

	testLiteralExpression(t, exp.Arguments[0], 1)
	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}

// TestCallExpressionParameterParsing
func TestCallExpressionParameterParsing(t *testing.T) {
	tests := []struct {
		input         string
		expectedIdent string
		expectedArgs  []string
	}{
		{
			input:         "add();",
			expectedIdent: "add",
			expectedArgs:  []string{},
		},
		{
			input:         "add(1);",
			expectedIdent: "add",
			expectedArgs:  []string{"1"},
		},
		{
			input:         "add(1, 2 * 3, 4 + 5);",
			expectedIdent: "add",
			expectedArgs:  []string{"1", "(2 * 3)", "(4 + 5)"},
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		require.True(t, ok)

		exp, ok := stmt.Expression.(*ast.CallExpression)
		require.True(t, ok)
		testIdentifier(t, exp.Function, tt.expectedIdent)

		require.Equal(t, len(exp.Arguments), len(tt.expectedArgs))

		for i, arg := range tt.expectedArgs {
			require.Equal(t, exp.Arguments[i].String(), arg)
		}
	}
}
