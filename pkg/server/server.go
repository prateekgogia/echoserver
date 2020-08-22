package server

import (
	"fmt"
	"log"
	"net"

	"github.com/prateekgogia/echoserver/api"

	"google.golang.org/grpc"
)

var (
	serverCertFile = "/usr/local/include/cert/server.crt"
	serverKeyFile  = "/usr/local/include/cert/server.key"
)

// Server represents the gRPC server
type Server struct {
	listener   net.Listener
	grpcServer *grpc.Server
	//
}

// New creates a server
func New() *Server {
	return &Server{}
}

// Run starts the server functionality
func (s *Server) Run(host string, port int) error {
	// create a net listener
	var err error
	if s.listener, err = startListener(host, port); err != nil {
		return err
	}
	// Create the TLS credentials
	// creds, err := credentials.NewServerTLSFromFile(serverCertFile, serverKeyFile)
	// if err != nil {
	// 	return fmt.Errorf("could not load TLS keys: %s", err)
	// }
	// create a gRPC server object
	// s.grpcServer = grpc.NewServer([]grpc.ServerOption{grpc.Creds(creds)}...)
	s.grpcServer = grpc.NewServer()

	// attach the Echo service to the server
	api.RegisterGraphRequestServer(s.grpcServer, s)

	// start the server, blocking call
	return s.grpcServer.Serve(s.listener)
}

// Stop the grpc server and listener
func (s *Server) Stop() error {
	defer log.Printf("Stopped gRPC echoserver")
	s.grpcServer.GracefulStop()
	return stopListener(s.listener)
}

func startListener(host string, port int) (net.Listener, error) {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}
	return l, err
}

func stopListener(lis net.Listener) error {
	return lis.Close()
}
