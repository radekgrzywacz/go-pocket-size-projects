package api

const (
	GameID         = "id"
	NewGameRoute   = "/games"
	GetStatusRoute = "/games/{" + GameID + "}"
	GuessRoute     = "/games/{" + GameID + "}"
)

type GuessRequest struct {
	Guess string `json:"guess"`
}

// GameResponse contains the information about a game.
type GameResponse struct {
	ID           string  `json:"id"`
	AttemptsLeft byte    `json:"attempts_left"`
	Guesses      []Guess `json:"guesses"`
	WordLength   byte    `json:"word_length"`
	Solution     string  `json:"solution,omitempty"`
	Status       string  `json:"status"`
}

// Guess is a pair of a word (submitted by the player) and its feedback.
type Guess struct {
	Word     string `json:"word"`
	Feedback string `json:"feedback"`
}
