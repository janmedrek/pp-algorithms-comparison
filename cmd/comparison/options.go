package main

import (
	"flag"
	"fmt"
)

type options struct {
	silent     bool
	hanoi      bool
	hanoiDisks int
}

func parseArgs() *options {
	silent := flag.Bool("silent", true, "Silent mode (moves are not logged).")
	hanoi := flag.Bool("hanoi", true, "Run Hanoi Towers comparison.")
	hanoiDisks := flag.Int("hanoiDisks", 5, "Number of disks in Hanoi Tower.")

	flag.Parse()

	return &options{
		silent:     *silent,
		hanoi:      *hanoi,
		hanoiDisks: *hanoiDisks,
	}
}

func (o *options) String() string {
	return fmt.Sprintf("--silent=%v --hanoi=%v --hanoiDisks=%d", o.silent, o.hanoiDisks, o.hanoiDisks)
}
