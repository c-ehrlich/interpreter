package main

import (
	"bufio"
	"bytes"
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 {
		filePath := os.Args[1]
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read file: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Running file %s\n", filePath)
		repl.Start(bufio.NewReader(bytes.NewReader(content)), os.Stdout, false)
		return
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout, true)
}
