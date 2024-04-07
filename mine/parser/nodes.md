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

# Parsing expressions

- `5 * 5 + 10` => (5 * 5) + 10 => 5 * 5 needs to be "deeper" in the tree
- parser needs to understand operator precedence, parens, etc
- tokens can appear in different positions with different meanings... `-5 - 10`
- parens have different meanings `5 * (add(2, 3) + 10)`
  - outer: "group expression"
  - inner: "call expression"
- essentially: meaning and validity of a token depend on context / the tokens that come before and after

- in monkey: everything except `let` and `return` is an expression
  - variables `5`, `foo`
  - prefix operators `!true`, `-5`
  - infix operators (also called binary operators)
    - arithmetic `5 + 5`
    - comparison operators `foo >= bar`
  - parens to group expressions and influence order of operation
  - call expressions `add(add(1, 2), add(3, 4))`
  - identifiers are expressions `foo`, `add(foo, bar)`
  - function literals are expressions `(fn(x) { return x }(5) + 10 ) * 10`
  - if statements are expressions `let result = if (foo > bar) { true } else { false }`
  - CONCLUSION: going linear wont work here

# Pratt Parsing (Operator Precedence)

- nud and led
  - lets the parser decide how to parse the "same" token depending on context
  - NUD: Null denotation - a method associated with a token that does not require anything to its left / can stand on its own... eg `-` in `-4`
    - for symbols, `nud()` usually returns the symbol itself, because they can stand alone
    - DOES NOT CARE ABOUT ANYTHING TO ITS LEFT
  - LED: Left denotation - a method associated with a token that comes after an expression ie `+` in `5 + 5`
    - DOES CARE ABOUT THINGS TO ITS LEFT
  - some operators have nud and led methods, eg `-`
- std: statement denotation ... tokens that can start a statement - `let`, `return`, etc
- lbp: left binding power ... determines precedence of tokens relative to each other ... `*` has higher lbp than `+`
- `find`... look in the current scope, then travel up the scope hierarchy. finally look in the symbol table.
  - new scopes need to know what exists in all previous scopes (usually through recursion) to know if an identifier is reserved or not

- `d (operator A) e (operator B) f`... is this `(d A e) B f` or `d A (e B f)`?
  - basically just this question over and over

- binding powers in js
  - 0: non binding operators like `;`
  - 10: assignment operators like `=`
  - 20: `?`
  - 30: `|| &&`
  - 40: relational operators like `===`
  - 50: `+ -`
  - 60: `* /`
  - 70: unary operators like `!`
  - 80: `. [ (`

# Pratt parsing again, this time from the book
- each token can have two parsing functions associated with it, depending on the token's position - infix or prefix
- prefix operator... "in front of" its operand... `++foo`
- postfix operator... "after" its operand... `foo++` - monkey doesn't have these
- infix operator... sits between to operands... `foo + bar`
- operator precedence... order of operations

- in our parser, every token's LBP and RBP are the same. this isn't true for every parser.
- if an operator should be right-associative (`(a + (b + c))`) it would need higher LBP than RBP