package server

import (
	"log"

	"github.com/prateekgogia/echoserver/api"
	"golang.org/x/net/context"
)

// EchoRequest generates response to a echo request
func (s *Server) EchoRequest(_ context.Context, in *api.EchoMessage) (*api.EchoMessage, error) {
	return echoRequest(in)
}

func echoRequest(in *api.EchoMessage) (*api.EchoMessage, error) {
	log.Printf("Receive message %s", in.Message)
	return &api.EchoMessage{Message: in.Message}, nil
}
