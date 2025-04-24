package main

import (
	"learngo/httpgordle/internal/handlers"
	"learngo/httpgordle/internal/repository"
	"net/http"
	"os"
)

func main() {
	db := repository.New()

	err := http.ListenAndServe(":8080", handlers.NewRouter(db))
	if err != nil {
		os.Exit(1)
	}
}
