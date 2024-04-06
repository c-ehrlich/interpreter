package ast

import (
	"monkey/token"
	"testing"
)

// `let myVar = anotherVar;`
func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	var expected = "let myVar = anotherVar;"
	if program.String() != expected {
		t.Errorf("program.String() wrong. expected=%q - got=%q",
			expected, program.String())
	}
}
