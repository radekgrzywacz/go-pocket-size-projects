package main

import (
	"learngo/httpgordle/internal/handlers"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(":8080", handlers.Mux())
	if err != nil {
		os.Exit(1)
	}
}
