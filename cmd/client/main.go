package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/prateekgogia/echoserver/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	msg  = flag.String("msg", "hello", "Message to be sent to the server")
	port = flag.Int("port", 8080, "gRPC server port")
)

func main() {
	flag.Parse()
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":%d", *port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewEchoClient(conn)
	response, err := c.EchoRequest(context.Background(), &api.EchoMessage{Message: *msg})
	if err != nil {
		log.Fatalf("Error when calling EchoRequest: %s", err)
	}
	log.Printf("Response from server: %s", response.Message)
}
