package server

import (
	"context"
	"fmt"
	"log"

	"github.com/prateekgogia/echoserver/api"
)

// EchoRequest generates response to a echo request
func (s *Server) EchoRequest(ctx context.Context, in *api.EchoMessage) (*api.EchoMessage, error) {
	var msg *api.EchoMessage
	var err error
	result := make(chan error)
	go func() {
		msg, err = echoRequest(in)
		result <- err
	}()
	select {
	case <-ctx.Done():
		log.Printf("Context Done err %v\n", ctx.Err())
		return nil, ctx.Err()
	case <-result:
		return msg, err
	}
}

func (s *Server) CreateGraph(ctx context.Context, graph *api.GraphObject) (*api.CreateGraphResponse, error) {
	fmt.Println("graph ", graph)
	return &api.CreateGraphResponse{GraphID: 1}, nil

}
func echoRequest(in *api.EchoMessage) (*api.EchoMessage, error) {
	log.Printf("Receive message %s", in.Message)
	return &api.EchoMessage{Message: in.Message}, nil
}
