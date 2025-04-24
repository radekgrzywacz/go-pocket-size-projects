package api

import "learngo/httpgordle/internal/session"

func ToGameResponse(g session.Game) GameResponse {
	apiGame := GameResponse{
		ID:           string(g.ID),
		AttemptsLeft: g.AttemptsLeft,
		Guesses:      make([]Guess, len(g.Guesses)),
		Status:       string(g.Status),
	}

	for index := 0; index < len(g.Guesses); index++ {
		apiGame.Guesses[index].Word = g.Guesses[index].Word
		apiGame.Guesses[index].Feedback = g.Guesses[index].Feedback
	}

	if g.AttemptsLeft == 0 {
		apiGame.Solution = ""
	}

	return apiGame
}
