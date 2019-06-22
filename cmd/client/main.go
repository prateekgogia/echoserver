package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/prateekgogia/echoserver/api"

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
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	conn, err = grpc.DialContext(ctx, fmt.Sprintf("%s:%d", host, *port),
		grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewEchoClient(conn)
	ctx, cancelReq := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelReq()
	response, err := c.EchoRequest(ctx, &api.EchoMessage{Message: *msg})
	if err != nil {
		log.Fatalf("Error when calling EchoRequest: %s", err)
	}
	log.Printf("Response from server: %s", response.Message)
}
