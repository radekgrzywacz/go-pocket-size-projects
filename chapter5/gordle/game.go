package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

func New(playerInput io.Reader, corpus []string, maxAttempts int) (*Game, error) {
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    []rune(strings.ToUpper(pickWord(corpus))),
		maxAttempts: maxAttempts,
	}

	return g, nil
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()
		fb := computeFeedback(guess, g.solution)

		fmt.Println(fb.String())
		if slices.Equal(guess, g.solution) {
			fmt.Printf("👏 You won! You found it in %d guess(es)", currentAttempt)
			return
		}

	}

	fmt.Printf(" 😞 You have lost... The solution was %s", string(g.solution))
}

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read the input")
			continue
		}

		guess := splitToUppercaseCharacters(string(playerInput))

		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle\n")
		} else {
			return guess
		}
	}
}

var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}

	return nil
}

func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

func computeFeedback(guess, solution []rune) feedback {
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error! Guess and solution are of different lengths!")
		return result
	}

	for posInGuess, character := range guess {
		if character == solution[posInGuess] {
			result[posInGuess] = correctPosition
			used[posInGuess] = true
		}
	}

	for posInGuess, character := range guess {
		if result[posInGuess] != absentCharacter {
			continue
		}

		for posInSolution, target := range solution {
			if used[posInSolution] {
				continue
			}
			if character == target {
				result[posInGuess] = wrongPosition
				used[posInSolution] = true
				break
			}
		}
	}
	return result
}
