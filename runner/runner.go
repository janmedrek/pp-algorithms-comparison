package runner

import (
	"log"
	"time"
)

type FunctionRunner interface {
	RunComparison(f1, f2 func(), label1, label2 string)
}

func NewFunctionRunner() FunctionRunner {
	return &functionRunner{}
}

type functionRunner struct {
}

func (r *functionRunner) RunComparison(f1, f2 func(), label1, label2 string) {
	f1Time := measureExecutionTime(f1, label1)
	f2Time := measureExecutionTime(f2, label2)

	printComparison(f1Time, f2Time, label1, label2)
}

func printComparison(f1, f2 time.Duration, label1, label2 string) {
	log.Printf("Function %s:", label1)
	log.Print("\t", f1)

	log.Printf("Function %s:", label2)
	log.Print("\t", f2)

	if f1 > f2 {
		diff := f1 - f2
		log.Printf("Function %s was slower and took %s longer to execute", label1, diff.String())
	} else {
		diff := f2 - f1
		log.Printf("Function %s was slower and took %s longer to execute", label2, diff.String())
	}
}

func measureExecutionTime(function func(), label string) time.Duration {
	log.Printf("Running func %s", label)

	timeStamp := time.Now()

	function()

	return time.Since(timeStamp)
}
