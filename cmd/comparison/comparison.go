package main

import (
	"github.com/janmedrek/pp-algorithms-comparison/runner"
	"github.com/janmedrek/pp-algorithms-comparison/towers"
)

func main() {
	opts := parseArgs()

	funcRunner := runner.NewFunctionRunner()

	if opts.hanoi {
		tower := towers.NewHanoiTower(opts.hanoiDisks, opts.silent)

		funcRunner.RunComparison(tower.SolveRecursive, tower.SolveIterative, "Hanoi-Recursive", "Hanoi-Iterative")
	}
}
