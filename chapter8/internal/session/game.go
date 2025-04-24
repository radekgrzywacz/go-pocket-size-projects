package session

import "errors"

type Game struct {
	ID           GameId
	AttemptsLeft byte
	Guesses      []Guess
	Status       Status
}

type GameId string

type Status string

const (
	StatusPlaying = "Playing"
	StatusWon     = "Won"
	StatusLost    = "Lost"
)

type Guess struct {
	Word     string
	Feedback string
}

var ErrGameOver = errors.New("Game over")
