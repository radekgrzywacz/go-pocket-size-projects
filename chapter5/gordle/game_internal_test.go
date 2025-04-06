package gordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGame_Ask(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 chars in english": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in arabic": {
			input: "مرحبا",
			want:  []rune("مرحبا"),
		},
		"5 characters in japanese": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		"3 characters in japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(strings.NewReader(tc.input), string(tc.want), len(tc.want))
			got := g.ask()

			if !slices.Equal(got, tc.want) {
				t.Errorf("gog = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGame_ValidateGuess(t *testing.T) {
	tc := map[string]struct {
		input    []rune
		expected error
	}{
		"valid input": {
			input:    []rune("GUESS"),
			expected: nil,
		},
		"Too long": {
			input:    []rune("LONGGUESS"),
			expected: errInvalidWordLength,
		},
		"Too short": {
			input:    []rune("GUE"),
			expected: errInvalidWordLength,
		},
		"Empty": {
			input:    []rune(nil),
			expected: errInvalidWordLength,
		},
	}

	for name, testCase := range tc {
		t.Run(name, func(t *testing.T) {
			g := New(nil, "SLICE", 0)

			err := g.validateGuess(testCase.input)
			if !errors.Is(err, testCase.expected) {
				t.Errorf("%c, expecte %q, got %q", testCase.input, testCase.expected, err)
			}
		})
	}
}

func TestGame_SplitToUppercaseCharacters(t *testing.T) {
	tt := map[string]struct {
		input    string
		expected []rune
	}{
		"Valid test": {
			input:    "thiss",
			expected: []rune("THISS"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			result := splitToUppercaseCharacters(tc.input)
			if !slices.Equal(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestFeedback_Equal(t *testing.T) {
	tt := map[string]struct {
		guess            string
		solution         string
		expectedFeedback feedback
	}{
		"nominal": {
			guess:            "HERTZ",
			solution:         "HERTZ",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double character": {
			guess:            "HELLO",
			solution:         "HELLO",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double character with wrong answer": {
			guess:            "HELLL",
			solution:         "HELLO",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, absentCharacter},
		},
		"five identical, but only two are there": {
			guess:            "LLLLL",
			solution:         "HELLO",
			expectedFeedback: feedback{absentCharacter, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
		"two identical, but not in the right position (from left to right)": {
			guess:            "HLLEO",
			solution:         "HELLO",
			expectedFeedback: feedback{correctPosition, wrongPosition, correctPosition, wrongPosition, correctPosition},
		},
		"three identical, but not in the right position (from left to right)": {
			guess:            "HLLLO",
			solution:         "HELLO",
			expectedFeedback: feedback{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"one correct, one incorrect, one absent (left of the correct)": {
			guess:            "LLLWW",
			solution:         "HELLO",
			expectedFeedback: feedback{wrongPosition, absentCharacter, correctPosition, absentCharacter, absentCharacter},
		},
		"swapped characters": {
			guess:            "HOLLE",
			solution:         "HELLO",
			expectedFeedback: feedback{correctPosition, wrongPosition, correctPosition, correctPosition, wrongPosition},
		},
		"absent character": {
			guess:            "HULFO",
			solution:         "HELFO",
			expectedFeedback: feedback{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"absent character and incorrect": {
			guess:            "HULPP",
			solution:         "HELPO",
			expectedFeedback: feedback{correctPosition, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			fb := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			if !tc.expectedFeedback.Equal(fb) {
				t.Errorf("Guess %q, got the wrong feedback, expected %q", fb, tc.expectedFeedback)
			}

		})
	}

}
