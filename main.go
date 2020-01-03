package main

import (
	"fmt"
	"monkey/repl"
	"os"
)

func main() {
	fmt.Println("Welcome to the Monkey programming language!")
	repl.Start(os.Stdin, os.Stdout)
}
