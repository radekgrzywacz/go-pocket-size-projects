package newgame

import (
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/session"
	"log"
	"net/http"
)

type gameAdder interface {
	Add(game session.Game)
}

func Handle(adder gameAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		game := createGame(adder)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		apiGame := api.ToGameResponse(game)
		err := json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			log.Printf("failed to write response: %s", err)
		}
	}
}

func createGame(db gameAdder) session.Game {
	return session.Game{}
}

func response(game session.Game) api.GameResponse {
	return api.GameResponse{}
}
