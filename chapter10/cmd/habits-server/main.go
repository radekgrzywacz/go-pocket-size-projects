package main

import (
	"fmt"
	"learngo/habits/internal/server"
	"learngo/habits/log"
	"os"
)

const port = 28710

func main() {
	lgr := log.New(os.Stdout)
	
	srv := server.New(lgr)
	
	err := srv.ListenAndServe(port)
	if err != nil {
	  fmt.Errorf("Error while running the server %s", err.Error())
	  os.Exit(1)
	}
	
}