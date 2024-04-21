# Enhance the interpreter
- [x] in addition to the repl, evaluate an entire file (`go run main.go foo.txt`)
- [ ] `void` type, especially for some builtin functions
- [ ] really understand how arrays work
- [ ] support `fn foo(bar) { return "baz" }`
- [ ] support utf-8 or unicode
- [x] support numbers formatted like `1_000_000`
- [ ] support floats
  - [ ] look for `.`, if there's a second one throw
- [ ] tofloat(int), toint(float..truncates), round(float), ceil(float), floor(float)
- [ ] support `<=` and `>=` (for numbers)
- [ ] support numbers formatted like `10e3` (just parse to int)
- [x] support strings
- [ ] support comments
  - [ ] `// foo`
  - [ ] `/* foo */`
- [ ] support more two char tokens, create `l.makeTwoCharToken` method
- [ ] add a lot more parser tests (ref has many, try own)
- [ ] support ternary like js
- [ ] support `++`, `--`
  - [ ] before number
  - [ ] after number - REQUIRES POSTFIX OPERATORS, WHICH MONKEY DOESN'T HAVE
- [ ] think of an operator that should be `(a OP (b OP c))` and break out LBP/RBP
  - [ ] ++ and -- might require this also...
- [ ] support `if (a && b)`, `if (a || b)`
- [ ] make `!0` evaluate to true instead of false
- [ ] do some dynamic type stuff (like js?)
  - [ ] maybe just type conversion operators?
- [ ] loops
- [ ] `else if`
- [ ] support `==` and `!=` for strings
- [ ] add more builtin functions
- [ ] dont allow assigning to the name of builtin functions
- [ ] add map, reduce, foreach... `map(arr, fn)` style

# Next steps
- [ ] the book states that parsers are usually generated instead of written. how can we generate a a monkey parser?
- [ ] add stack traces to the error object