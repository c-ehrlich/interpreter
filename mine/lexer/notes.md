# Lexer

Source code => tokens => AST

Lexer does code => tokens

source code: `let x = 5 + 5;`
tokens: [ LET, IDENTIFIER('x'), EQUAL_SIGN, INTEGER(5), PLUS_SIGN, INTEGER(5), SEMICOLON ]

"real" lexer: also might attach filename, line, column to each token - for errors etc

Lexer goals
- go through input, and output the next token
- dont buffer or save tokens, just run `NextToken()` each time
- We use `string` for the text, real lexer might use `io.Reader` and the filename