package repository

import (
	"fmt"
	"learngo/httpgordle/internal/session"
)

type GameRepository struct {
	storage map[session.GameId]session.Game
}

func New() *GameRepository {
	return &GameRepository{
		storage: make(map[session.GameId]session.Game),
	}
}

func (gr *GameRepository) Add(game session.Game) error {
	_, ok := gr.storage[game.ID]
	if ok {
		return fmt.Errorf("gammeId %s already exists", game.ID)
	}

	gr.storage[game.ID] = game

	return nil
}

func (gr *GameRepository) Find(game session.Game) (session.Game, error) {
	game, ok := gr.storage[game.ID]
	if !ok {
		return session.Game{}, fmt.Errorf("No game with id %s", game.ID)
	}

	return game, nil
}

func (gr *GameRepository) Update(game session.Game) error {
	_, ok := gr.storage[game.ID]
	if !ok {
		return fmt.Errorf("No game with id %s", game.ID)
	}
	gr.storage[game.ID] = game
	return nil
}
