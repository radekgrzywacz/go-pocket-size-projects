package newgame

import (
	"learngo/httpgordle/internal/api"
	"net/http"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	if req.Method != api.NewGameMethod {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Creating a new game"))
}
