package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/conao3/go-glisp/lexer"
)

const PROMPT = "glisp> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			fmt.Printf("\n")
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		fmt.Printf("%+v\n", l)
	}
}
