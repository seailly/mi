package evaluator

import (
	"testing"

	"github.com/seailly/mi/lexer"
	"github.com/seailly/mi/object"
	"github.com/seailly/mi/parser"
	"github.com/stretchr/testify/require"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct{
		input string
		expected int64
	}{
		{ "5", 5 },
		{ "10", 10 },
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) {
	result, ok := obj.(*object.Integer)
	require.True(t, ok)
	require.Equal(t, expected, result.Value)
}