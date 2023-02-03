package main

import (
	"fmt"
	"os"

	"github.com/conao3/go-glisp/repl"
)

func main() {
	fmt.Println("Hello, playground")
	repl.Start(os.Stdin, os.Stdout)
}
