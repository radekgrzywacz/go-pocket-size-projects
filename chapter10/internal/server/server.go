package server

import (
	"fmt"
	"learngo/habits/api"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type Server struct {
	api.UnimplementedHabitsServer
	lgr Logger
}

func New(lgr Logger) *Server {
	return &Server{
		lgr: lgr,
	}
}

func (s *Server) ListenAndServe(port int) error {
	const addr = "127.0.0.1"
	
	listener, err := net.Listen("tcp", net.JoinHostPort(addr, strconv.Itoa(port)))
	if err != nil {
	  return fmt.Errorf("Unable to listen to tcp port %d: %w", port, err)
	}
	
	grpcServer := grpc.NewServer()
	api.RegisterHabitsServer(grpcServer, s)
	
	s.lgr.Logf("Starting server on port %d\n", port)
	
	err = grpcServer.Serve(listener)
	if err != nil {
	  return fmt.Errorf("error while listening: %w", err)
	}
	return nil
}

type Logger interface {
	Logf(format string, args ...any)
}