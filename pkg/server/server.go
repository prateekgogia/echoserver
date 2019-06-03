package server

import (
	"fmt"
	"log"
	"net"

	"github.com/prateekgogia/echoserver/api"

	"google.golang.org/grpc"
)

// Server represents the gRPC server
type Server struct {
	listener   net.Listener
	grpcServer *grpc.Server
}

// New creates a server
func New() *Server {
	return &Server{}
}

// Run starts the server functionality
func (s *Server) Run(port int) error {
	// create a net listener
	var err error
	if s.listener, err = startListener(port); err != nil {
		return err
	}
	// create a gRPC server object
	s.grpcServer = grpc.NewServer()

	// attach the Echo service to the server
	api.RegisterEchoServer(s.grpcServer, s)

	// start the server, blocking call
	return s.grpcServer.Serve(s.listener)
}

// Stop the grpc server and listener
func (s *Server) Stop() error {
	defer log.Printf("Stopped gRPC echoserver")
	s.grpcServer.GracefulStop()
	return stopListener(s.listener)
}

func startListener(port int) (net.Listener, error) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}
	return l, err
}

func stopListener(lis net.Listener) error {
	return lis.Close()
}
