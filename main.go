package main

import (
	"flag"
	"io"
	"os"

	"github.com/conao3/go-glisp/repl"
)

func main() {
	var (
		stageReader    bool
		stageEvaluator bool
		Input          string
	)
	flag.BoolVar(&stageReader, "sr", false, "reader stage")
	flag.BoolVar(&stageEvaluator, "se", false, "evaluator stage")
	flag.StringVar(&Input, "i", "", "input file")

	flag.Parse()

	stage := repl.StageDefault

	if stageReader {
		stage = repl.StageReader
	} else if stageEvaluator {
		stage = repl.StageEvaluator
	}

	var inpt io.Reader
	switch Input {
	case "":
		inpt = os.Stdin
	case "-":
		inpt = os.Stdin
	default:
		file, err := os.Open(Input)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		inpt = file
	}
	repl.Start(inpt, os.Stdout, stage)
}
