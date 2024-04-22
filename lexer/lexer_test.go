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
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q, %q",
				i, tt.expectedType, tok.Type, tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q, %q",
				i, tt.expectedLiteral, tok.Literal, tok.Type)
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
	input := `!-*/5;
5 < 10 > 5;`

	tests := []TokenTest{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
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
"foo bar"
"foo\"bar"`

	tests := []TokenTest{
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.STRING, `foo"bar`},
		{token.EOF, ""},
	}

	testTokens(t, input, tests)
}

func TestNextTokenNumbers(t *testing.T) {
	input := `1 123 1_000_000`

	tests := []TokenTest{
		{token.INT, "1"},
		{token.INT, "123"},
		{token.INT, "1000000"},
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

func TestNextTokenColon(t *testing.T) {
	input := `{"foo": "bar"}`

	tests := []TokenTest{
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	testTokens(t, input, tests)
}

func TestSingleLineComment(t *testing.T) {
	input := `5; // foo
4;
// bar
3;`

	tests := []TokenTest{
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
	}

	testTokens(t, input, tests)
}

func TestMultiLineComment(t *testing.T) {
	input := `5; /* foo
 * bar
 */ 4;
3; /* baz */ 2;`

	tests := []TokenTest{
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
	}

	testTokens(t, input, tests)
}

func TestFloat(t *testing.T) {
	input := `123.456;
.5;
1.;
1_2_3.4_5_6;
123.456.789;`

	tests := []TokenTest{
		{token.FLOAT, "123.456"},
		{token.SEMICOLON, ";"},
		{token.FLOAT, ".5"},
		{token.SEMICOLON, ";"},
		{token.FLOAT, "1."},
		{token.SEMICOLON, ";"},
		{token.FLOAT, "123.456"},
		{token.SEMICOLON, ";"},
		{token.ILLEGAL, "123.456.789"},
		{token.SEMICOLON, ";"},
	}

	testTokens(t, input, tests)
}

func TestIncrementDecrement(t *testing.T) {
	input := `foo = ++bar;
--baz;`

	tests := []TokenTest{
		{token.IDENT, "foo"},
		{token.ASSIGN, "="},
		{token.INCREMENT, "++"},
		{token.IDENT, "bar"},
		{token.SEMICOLON, ";"},
		{token.DECREMENT, "--"},
		{token.IDENT, "baz"},
		{token.SEMICOLON, ";"},
	}

	testTokens(t, input, tests)
}

func TestLteGte(t *testing.T) {
	input := `foo <= bar;
baz >= biz;`

	tests := []TokenTest{
		{token.IDENT, "foo"},
		{token.LTE, "<="},
		{token.IDENT, "bar"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "baz"},
		{token.GTE, ">="},
		{token.IDENT, "biz"},
		{token.SEMICOLON, ";"},
	}

	testTokens(t, input, tests)
}

func TestAndOr(t *testing.T) {
	input := `(foo && bar) || (baz || biz);`

	tests := []TokenTest{
		{token.LPAREN, "("},
		{token.IDENT, "foo"},
		{token.AND, "&&"},
		{token.IDENT, "bar"},
		{token.RPAREN, ")"},
		{token.OR, "||"},
		{token.LPAREN, "("},
		{token.IDENT, "baz"},
		{token.OR, "||"},
		{token.IDENT, "biz"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	testTokens(t, input, tests)
}

func TestForLoop(t *testing.T) {
	input := `while (i < 10) { if (i == 5) { return i }; ++i; }`

	tests := []TokenTest{
		{token.WHILE, "while"},
		{token.LPAREN, "("},
		{token.IDENT, "i"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "i"},
		{token.EQ, "=="},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "i"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.INCREMENT, "++"},
		{token.IDENT, "i"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	testTokens(t, input, tests)
}
