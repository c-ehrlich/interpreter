package lexer

import (
	"testing"

	"monkey/token"
)

type TokenTest struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func testTokens(
	t *testing.T,
	input string,
	tests []TokenTest,
) {
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenSingleChars(t *testing.T) {
	input := `=+(){},;`

	tests := []TokenTest{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	testTokens(t, input, tests)
}

func TestNextTokenComplex(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
`
	tests := []TokenTest{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	testTokens(t, input, tests)
}

func TestNextTokenMoreTokens(t *testing.T) {
	input := `!-/*5;
5 < 10 > 5;`

	tests := []TokenTest{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	testTokens(t, input, tests)
}

func TestNextTokenTwoCharTokens(t *testing.T) {
	input := `10 == 10;
9 != 10;`

	// "slice of structs"
	tests := []TokenTest{
		// if you use the same order as the struct, you dont need to specify the keys
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "9"},
		{token.NOT_EQ, "!="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
	}

	testTokens(t, input, tests)
}

func TestNextTokenStrings(t *testing.T) {
	input := `"foobar"
"foo bar"`

	tests := []TokenTest{
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.EOF, ""},
	}

	testTokens(t, input, tests)
}

func TestNextTokenArrayBrackets(t *testing.T) {
	input := `[1, 2];`

	tests := []TokenTest{
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	testTokens(t, input, tests)
}
