# Enhance the interpreter
- [ ] support `fn foo(bar: string) { return "baz" }`
- [ ] support utf-8 or unicode
- [ ] support numbers formatted like `1_000_000`
- [ ] support other number types (float, hex, octal)
- [ ] support `<=` and `>=`
- [ ] support strings
- [ ] support comments
- [ ] support more two char tokens, create `l.makeTwoCharToken` method
- [ ] add a lot more parser tests (ref has many, try own)
- [ ] support ternary like js
- [ ] support `++`, `--` - REQUIRES POSTFIX OPERATORS, WHICH MONKEY DOESN'T HAVE

# Next steps
- [ ] the book states that parsers are usually generated instead of written. how can we generate a a monkey parser?