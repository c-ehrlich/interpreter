# Parser
- take input data (usually text) and turn it into a data structure
- example: JSON.parse()
- usually preceded by a lexer (JSON.parse includes the lexing step)
- With JSON.parse the data structure is human readable... with language parsers usually not
- usually parse into AST (abstract syntax tree)
  - abstract: some parts of the source code are omitted (semicolon, newlines, comments, parens, etc - depending on the language)
- most ASTs are very similar, but there is no "one true AST format" - implementations differ in details

- there are a lot of parser-generators - yacc, bison, ANTLR
  - they use "context-free grammar" (CFG)... a set of rules for how to form correct sentences in a language
    - most common: Backus-Naur Form (BNF), Extended BNF (EBNF)
- parsers are mostly solved
- you probably shouldn't write your own parser except for educational purposes

- two main strategies for parsing: top-down, bottom-up
  - some variations like "recursive descent parsing", "early parsing", "predictive parsing"

- we will write a recursive descent parser - specifically "top down operator precedence", or "Pratt parser"
- top down starts with the root node of AST

- concessions for ours
  - won't be fast
  - won't have formal proof of correctness
  - error-recovery and detection of erroneous syntax won't be great (this requires a lot of study of theory!)

```
let x = 10;
let y = 15;
let add = fn(a, b) {
    return a + b;
};
```
- `x`, `y`, `add`: identifiers
- `10`, `fn(a, b){return a + b}`: expressions

- `let <identifier> = <expression>`
- `let x = 10` ... `10`: literal expression
- `let sum = add(x+y)` ... `add(x+y)`: function expression

- statement vs expression: expressions produce values, statements do not
- `return 5` doesn't produce a value, `add(5, 5)` does
- different things are expressions in different languages

- to parse `5 + 5`, we need to parse `5 +` and then call `parseExpression` again on the right side
- ...since the right side might be `add(1, 2)`

- parser repeatedly advances the tokens and checks the current token to decide what to do next
  - call a parsing function
  - throw an error