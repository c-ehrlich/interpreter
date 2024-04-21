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
1.0 / 2.0; // 0.5
toint(1.0); // 1
toint(1.9); // 1
tofloat(1); // 1.0
toint(floatA) == intA; // true
ceil(floatA); // 2.0
floor(floatA); // 1.0
floor(-1.1); // -2.0
round(1.5); // 2.0
1.2 > 1.1; // true
// etc
```

### Underscores in numbers
```js
let foo = 1_000_000; 
```

### Increment and decrement operators
```js
let foo = ++1; // 2
let bar = --foo; // 1
```

### GTE and LTE
```js
2 >= 1; // true
1.5 >= 1.5; // true
1 >= 2; // false

2 <= 1; // false
1.5 <= 1.5; // false
1 <= 2; // false
```

### String comparison
```js
"apples" == "apples"; // true
"apples" == "oranges"; // false
"apples" != "apples"; // false
"apples" != "oranges"; // true

### Comments
```js
let a = 5; // foo
// bar
let b = 10;
```