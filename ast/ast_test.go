package ast

import (
	"testing"

	"github.com/seailly/mi/token"
)

// TestString
func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&MutStatement{
				Token: token.Token{
					Type:    token.MUT,
					Literal: "mut",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "mut myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got %q", program.String())
	}
}
