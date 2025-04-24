package guess

import (
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/session"
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue(api.GameID)
	if id == "" {
		http.Error(w, "missing the id of the game", http.StatusBadRequest)
		return
	}

	// Read the request, containing the guess, from the body of the input.
	r := api.GuessRequest{}
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	game := guess(id)

	apiGame := api.ToGameResponse(game)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(apiGame)
	if err != nil {
		// The header has already been set. Nothing much we can do here.
		log.Printf("failed to write response: %s", err)
	}
}

func guess(id string) session.Game {
	return session.Game{
		ID: session.GameId(id),
	}
}
