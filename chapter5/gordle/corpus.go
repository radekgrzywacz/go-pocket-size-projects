package gordle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const ErrCorpusIsEmpty = corpusError("Corpus is empty")

func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to oper %q for reading: %w", path, err)
	}

	if len(data) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	words := strings.Fields(string(data))
	return words, nil
}

func pickWord(corpus []string) string {
	index := rand.Intn(len(corpus))
	return corpus[index]
}
