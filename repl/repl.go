package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/conao3/go-glisp/token"
)

const PROMPT = "glisp> "

type Stage int

const (
	StageDefault Stage = iota
	StageTokenizer
	StageLexer
	StageReader
	StageEvaluator
)

func Start(in io.Reader, out io.Writer, stage Stage) {
	if stage == StageDefault {
		stage = StageTokenizer
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
		t := token.New(line)

		if stage == StageTokenizer {
			fmt.Printf("%+v\n", t)
		}

		if stage == StageLexer || stage == StageReader || stage == StageEvaluator {
			panic("not implemented")
		}
	}
}
