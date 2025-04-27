package main

import (
	"fmt"
	solver "learngo/maze_solver/solver/internal"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	s, err := solver.New(inputFile)
	if err != nil {
		exit(err)
	}

	err = s.Solve()
	if err != nil {
		exit(err)
	}

	err = s.SaveSolution(outputFile)
	if err != nil {
		exit(err)
	}
}

func usage() {
	_, _ = fmt.Fprintln(os.Stderr, "Usage: maze_solver input.png output.png")
	os.Exit(1)
}

func exit(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %s", err)
	os.Exit(1)
}
