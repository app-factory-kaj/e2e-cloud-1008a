package handlers

import (
	"context"

	"greeter/internal/gen"
)

// Server implements the greeter API.
type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetGreeting(_ context.Context, request gen.GetGreetingRequestObject) (gen.GetGreetingResponseObject, error) {
	name := request.Params.Name
	if name == "" {
		name = "World"
	}
	return gen.GetGreeting200JSONResponse{
		Message: "Hello, " + name + "!",
	}, nil
}
