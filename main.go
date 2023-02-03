package main

import (
	"flag"
	"os"

	"github.com/conao3/go-glisp/repl"
)

func main() {
	var (
		stageTokenizer bool
		stageLexer     bool
		stageReader    bool
		stageEvaluator bool
	)
	flag.BoolVar(&stageTokenizer, "st", false, "tokenizer stage")
	flag.BoolVar(&stageLexer, "sl", false, "lexer stage")
	flag.BoolVar(&stageReader, "sr", false, "reader stage")
	flag.BoolVar(&stageEvaluator, "se", false, "evaluator stage")

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

	repl.Start(os.Stdin, os.Stdout, stage)
}
