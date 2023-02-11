package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/conao3/go-glisp/evaluator"
	"github.com/conao3/go-glisp/reader"
	"github.com/conao3/go-glisp/types"
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
	env := types.NewEnvironment()

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

		res := evaluator.Eval(exp, env)

		if stage == StageEvaluator {
			fmt.Printf("%+v\n", res)
		}
	}
}
