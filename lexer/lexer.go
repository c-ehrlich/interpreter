package lexer

import (
	"monkey/token"
	"strings"
)

type Lexer struct {
	// the entire input to be tokenized
	input string
	// current position in input (points to current char)
	position int
	// current reading position in input (after current char)
	readPosition int
	// current character under examination
	// we're only going to support ASCII characters so we use byte
	// to support unicode or utf-8 we'd use rune
	ch byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Give us the next character and advance our position in the input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL" character - either we haven't read anything yet or we're at the end of the input
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	// operators
	case '=':
		if l.peekChar() == '=' {
			tok = readTwoCharacterToken(l, token.EQ)
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		if l.peekChar() == '+' {
			tok = readTwoCharacterToken(l, token.INCREMENT)
		} else {
			tok = newToken(token.PLUS, l.ch)
		}
	case '-':
		if l.peekChar() == '-' {
			tok = readTwoCharacterToken(l, token.DECREMENT)
		} else {
			tok = newToken(token.MINUS, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			tok = readTwoCharacterToken(l, token.NOT_EQ)
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		if l.peekChar() == '/' {
			l.skipComment()
			return l.NextToken()
		} else if l.peekChar() == '*' {
			l.skipMultilineComment()
			return l.NextToken()
		} else {
			tok = newToken(token.SLASH, l.ch)
		}
	case '%':
		tok = newToken(token.MODULO, l.ch)
	case '<':
		if l.peekChar() == '=' {
			tok = readTwoCharacterToken(l, token.LTE)
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			tok = readTwoCharacterToken(l, token.GTE)
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case '&':
		if l.peekChar() == '&' {
			tok = readTwoCharacterToken(l, token.AND)
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case '|':
		if l.peekChar() == '|' {
			tok = readTwoCharacterToken(l, token.OR)
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	// delimiters
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	// strings
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	// numbers, identifiers, illegal
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigitOrDecimalPoint(l.ch) {
			tok = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position] // range of characters that form the identifier [position, l.position]
}

func (l *Lexer) readNumber() token.Token {
	isFloat := false
	multipleDecimals := false
	position := l.position
	for isNumberComponent(l.ch) {
		if l.ch == '.' {
			if isFloat {
				multipleDecimals = true
			}
			isFloat = true
		}
		l.readChar()
	}
	numString := l.input[position:l.position]
	withoutUnderscores := strings.ReplaceAll(numString, "_", "")
	if multipleDecimals {
		return token.Token{Type: token.ILLEGAL, Literal: withoutUnderscores}
	}
	if isFloat {
		return token.Token{Type: token.FLOAT, Literal: withoutUnderscores}
	}
	return token.Token{Type: token.INT, Literal: withoutUnderscores}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	l.skipWhitespace()
}

func (l *Lexer) skipMultilineComment() {
	for l.ch != '*' || l.peekChar() != '/' {
		l.readChar()
	}
	l.readChar()
	l.readChar()
	l.skipWhitespace()
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readString() string {
	str := ""
	l.readChar() // skip initial `"`

	for {
		// handle end of string
		if l.ch == '"' || l.ch == 0 {
			// l.readChar() // skip final `"`
			break
		}

		// handle escaped quotes
		if l.ch == '\\' && l.peekChar() == '"' {
			str += "\""
			l.readChar()
			l.readChar()
			continue
		}

		// handle regular characters
		str += string(l.ch)
		l.readChar()
	}

	return str
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// letters are a-z, A-Z, and _
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigitOrDecimalPoint(ch byte) bool {
	return ('0' <= ch && ch <= '9') || ch == '.'
}

func isNumberComponent(ch byte) bool {
	return isDigitOrDecimalPoint(ch) || ch == '_' || ch == '.'
}

func readTwoCharacterToken(l *Lexer, tokenType token.TokenType) token.Token {
	ch := l.ch
	l.readChar()
	literal := string(ch) + string(l.ch)
	return token.Token{Type: tokenType, Literal: literal}
}
