# Interpreter

## Development

1. Install gow `go install github.com/mitranim/gow@latest`
2. Run tests in watch mode `gow test ./...`
3. To start the repl: `go run main.go`
4. To pass a file: `go run main.go program.monkey`

## Additional Features

### Pass a file to the repl
```
go run main.go program.monkey
```

### Floating point numbers
```js
let floatA = 1.1;
let floatB = 1.2;
let intA = 1;
let intB = 2;
1.0 / 2.0; // 0.500000
toint(floatA) == intA; // true
ceil(floatA); // 2.00000
floor(floatA); // 1.00000
floor(-1.1); // -2.00000
round(1.5); // 2.00000
1.2 > 1.1; // true
// etc
```

### Underscores in numbers
```js
let foo = 1_000_000; 
```

### Comments
```js
let a = 5; // foo
// bar
let b = 10;
```