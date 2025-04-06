package main

import (
	"fmt"
	"gordle/gordle"
	"os"
)

const maxAttempts = 6

func main() {
	corpus, err := gordle.ReadCorpus("gordle/corpus/english.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to read corpus: %s", err)
	}
	g, err := gordle.New(os.Stdin, corpus, maxAttempts)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to start corpus: %s", err)
		return
	}
	g.Play()
}
