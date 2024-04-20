# Enhance the interpreter
- [ ] really understand how arrays work
- [ ] support `fn foo(bar) { return "baz" }`
- [ ] support utf-8 or unicode
- [ ] support numbers formatted like `1_000_000`
- [ ] support other number types (float, hex, octal?)
- [ ] support `<=` and `>=`
- [ ] support strings
- [ ] support comments
- [ ] support more two char tokens, create `l.makeTwoCharToken` method
- [ ] add a lot more parser tests (ref has many, try own)
- [ ] support ternary like js
- [ ] support `++`, `--` - REQUIRES POSTFIX OPERATORS, WHICH MONKEY DOESN'T HAVE
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

# Next steps
- [ ] the book states that parsers are usually generated instead of written. how can we generate a a monkey parser?
- [ ] add stack traces to the error object