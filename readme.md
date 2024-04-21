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

### Underscores in numbers
```
let foo = 1_000_000; 
```

### Comments
```
let a = 5; // foo
// bar
let b = 10;
```