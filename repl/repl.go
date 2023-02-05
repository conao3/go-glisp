package repl

import (
	"bufio"
	"fmt"
	"github.com/conao3/go-glisp/reader"
	"io"
)

const PROMPT = "glisp> "

type Stage int

const (
	StageDefault Stage = iota
	StageReader
	StageEvaluator
)

func Start(in io.Reader, out io.Writer, stage Stage) {
	if stage == StageDefault {
		stage = StageReader
	}

	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			fmt.Printf("\n")
			return
		}

		line := scanner.Text()
		r := reader.New(line)
		exp := r.Read()

		if stage == StageReader {
			fmt.Printf("%+v\n", exp)
		}

		if stage == StageEvaluator {
			panic("not implemented")
		}
	}
}
