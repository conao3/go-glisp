package main

import (
	"flag"
	"io"
	"os"

	"github.com/conao3/go-glisp/repl"
)

func main() {
	var (
		stageTokenizer bool
		stageLexer     bool
		stageReader    bool
		stageEvaluator bool
		Input 		string
	)
	flag.BoolVar(&stageTokenizer, "st", false, "tokenizer stage")
	flag.BoolVar(&stageLexer, "sl", false, "lexer stage")
	flag.BoolVar(&stageReader, "sr", false, "reader stage")
	flag.BoolVar(&stageEvaluator, "se", false, "evaluator stage")
	flag.StringVar(&Input, "i", "", "input file")

	flag.Parse()

	stage := repl.StageDefault

	if stageTokenizer {
		stage = repl.StageTokenizer
	} else if stageLexer {
		stage = repl.StageLexer
	} else if stageReader {
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
