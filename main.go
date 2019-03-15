package main

import (
	"fmt"
	"github.com/cosmonawt/monkey/repl"
	"os"
)

func main() {
	fmt.Println("Welcome to the Monkey programming language!")
	repl.Start(os.Stdin, os.Stdout)
}
