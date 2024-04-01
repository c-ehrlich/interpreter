package token

/**
 * Most "real" tokenisers would use something more efficient
 * than strings to represent tokens.
 */
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // Unknown token
	EOF     = "EOF"     // Parser can stop

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  "FUNCTION",
	"let": "LET",
}

// check if the identifier is a keyword
func LookupIdent(ident string) TokenType {
	if tokenType, found := keywords[ident]; found {
		return tokenType
	}
	return IDENT
}
