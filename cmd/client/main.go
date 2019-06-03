package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/prateekgogia/echoserver/api"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	msg            = flag.String("msg", "hello", "Message to be sent to the server")
	port           = flag.Int("port", 8080, "gRPC server port")
	host           = "localhost"
	serverCertFile = "/usr/local/include/cert/server.crt"
)

func main() {
	flag.Parse()
	var conn *grpc.ClientConn
	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile(serverCertFile, "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}
	conn, err = grpc.Dial(fmt.Sprintf("%s:%d", host, *port),
		grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewEchoClient(conn)
	response, err := c.EchoRequest(context.Background(),
		&api.EchoMessage{Message: *msg})
	if err != nil {
		log.Fatalf("Error when calling EchoRequest: %s", err)
	}
	log.Printf("Response from server: %s", response.Message)
}
