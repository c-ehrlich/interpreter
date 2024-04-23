# Interpreter

## Example

### 1. Write a program
```js
// program.monkey
let fizz = fn(x){
    if (x % 15 == 0) { return "fizzbuzz"; }
    if (x % 3 == 0) { return "fizz"; }
    if (x % 5 == 0) { return "buzz"; }
    return x;
};

let i = 1;

while (i <= 100) {
    puts(fizz(i))
    ++i;
}
```

### 2. Run the program

`go run main.go program.monkey`

### 3. Be amazed

```
1
2
fizz
4
buzz
// etc...
```

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

### Modulo
```js
3 % 2; // 1
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
```

### Logical operators
```js
true && true; // true
true && false; // false
true || false; // true
false || false; // true
// they can compare booleans, integers, floats, strings,
// and any expression that evaluates to one of these types
1 && 2; // true
1 && 0; // false
1.5 && 0.0; // false
"foo" && ""; // false
0 || "foo"; // true
if (true && false) { "true" } else { "false" }; // "false"
```

### Loops
```js
let i = 0;
while (i < 10) {
    if (i == 5) {
        return i;
    }
}
```

#### And they treat lexical scope correctly!
```js
let i = 0;
let j = 0;
while (i < 10) {
    let j = 5;
    ++i;
}
puts(j); // 0
```

### Escape quotes in strings
```js
"foo\"bar"; // behind the scenes: `foo\"bar`, prints as `foo"bar`
```

### Single Line Comments
```js
let a = 5; // foo
// bar
let b = 10;
```

### Multi Line Comments
```js
/**
 * foo
 */ let c = 15;
let d = 20; /* bar */ let e = 25
```
