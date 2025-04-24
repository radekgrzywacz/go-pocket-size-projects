package getstatus

import (
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/session"
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, request *http.Request) {
	id := request.PathValue(api.GameID)
	if id == "" {
		http.Error(w, "missing the id of the game", http.StatusBadRequest)
		return
	}

	game := getGame(id)
	apiGame := api.ToGameResponse(game)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
}

func getGame(id string) session.Game {
	return session.Game{
		ID: session.GameId(id),
	}
}
