package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/prateekgogia/echoserver/api"

	"google.golang.org/grpc"
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
	// creds, err := credentials.NewClientTLSFromFile(serverCertFile, "")
	// if err != nil {
	// 	log.Fatalf("could not load tls cert: %s", err)
	// }
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", host, *port),
		grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewGraphRequestClient(conn)
	ctx, cancelReq := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelReq()

	response, err := c.CreateGraph(ctx, createGraphObject())
	if err != nil {
		log.Fatalf("Error when calling EchoRequest: %s", err)
	}
	log.Printf("Response from server: Graph ID is %d", response.GraphID)
}

func createGraphObject() *api.GraphObject {
	result := &api.GraphObject{}
	result.AdjacencyList = make(map[int32]*api.Neighbors)
	result.AdjacencyList[1] = &api.Neighbors{Nodes: []int32{2, 3}}
	result.AdjacencyList[2] = &api.Neighbors{Nodes: []int32{1, 4}}
	result.AdjacencyList[3] = &api.Neighbors{Nodes: []int32{1, 2}}
	result.AdjacencyList[4] = &api.Neighbors{Nodes: []int32{2}}
	return result
}

// 1 [2, 3]
// 2 [1, 4]
// 3 [1, 2]
// 4 [2]
