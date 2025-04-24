package handlers

import (
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/handlers/getstatus"
	"learngo/httpgordle/internal/handlers/guess"
	"learngo/httpgordle/internal/handlers/newgame"
	"learngo/httpgordle/internal/repository"
	"net/http"
)

// NewRouter returns a router that listens for request to the following endpoints:
//   - Create a new game
func NewRouter(db *repository.GameRepository) *http.ServeMux {
	r := http.NewServeMux()

	// Register each endpoint.
	r.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handle(db))
	r.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handle(db))
	r.HandleFunc(http.MethodPut+" "+api.GuessRoute, guess.Handle(db))

	return r
}
