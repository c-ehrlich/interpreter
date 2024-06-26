# Enhance the interpreter
- [x] in addition to the repl, evaluate an entire file (`go run main.go foo.txt`)
- [x] support numbers formatted like `1_000_000`
- [x] support floats
  - [x] look for `.`, if there's a second one throw
  - [x] negative
  - [x] negative test
  - [x] arithmetic
  - [x] arithmetic tests
  - [x] boolean logic
  - [x] boolean logic tests
- [x] support strings
- [x] support comments `// foo`, `let a = 1; // foo`
- [x] support `++`, `--` before number
- [x] `toint` function
- [x] `tofloat` function
- [x] `ceil` function
- [x] `floor` function
- [x] `round` function
- [x] extract `readTwoCharacterToken`
- [x] support `<=` and `>=` (ints)
- [x] `bool(anything)`
- [x] support `if (a && b)`, `if (a || b)`
- [x] support block comments `/**\n * foo\n */`
- [x] escape `\"` in strings
- [x] support `==` and `!=` for strings
- [x] while loop!
- [x] modulo
- [x] correctly handle lexical environment in while loop
- [ ] standardize how prefix expressions are handled
  - [ ] pass the whole node, grab the value downstream
  - [ ] figure out if/how this impacts error handling
  - [ ] should be as close to the elegance of the previous system as possible
- [ ] `let i = 0; i = 1;`
- [ ] `else if`
- [ ] treesitter https://github.com/jamestrew/tree-sitter-monkey
- [ ] support `10e3` notation (just parse to int)
- [ ] add map, reduce, foreach... php style `map(arr, fn)`
- [ ] make `bool(0)` evaluate to false
- [ ] make `!0` evaluate to true instead of false
- [ ] `void` type, especially for some builtin functions?
- [ ] alternative function syntax `fn foo(bar) { return "baz" }`
- [ ] support utf-8 or unicode
- [ ] support ternary like js `let foo = bar ? "baz" : "biz"`
- [ ] support `++`, `--` after number - REQUIRES POSTFIX OPERATORS
- [ ] think of an operator that should be `(a OP (b OP c))` and break out LBP/RBP
  - [ ] ++ and -- might require this also...
- [ ] dont allow assigning to the name of builtin functions

# Next steps
- [ ] the book states that parsers are usually generated instead of written. how can we generate a monkey parser?
- [ ] add stack traces to the error object