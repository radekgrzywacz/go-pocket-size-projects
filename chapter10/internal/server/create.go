package server

import (
	"context"
	"learngo/habits/api"
)

func (s *Server) CreateHabit(
	_ context.Context,
	request *api.CreateHabitRequest,
) (*api.CreateHabitResponse, error) {
	s.lgr.Logf("CreateHabit request received: %s", request)
	
	return &api.CreateHabitResponse{
		Habit: &api.Habit{},
	}, nil
}