package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"
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

	log.Println(" Send Echo Request Response ")
	c := api.NewEchoClient(conn)
	ctx, cancelReq := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelReq()
	response, err := c.EchoRequest(ctx, &api.EchoMessage{Message: *msg})
	if err != nil {
		log.Fatalf("Error when calling EchoRequest: %s", err)
	}
	log.Printf("Response from server: %s", response.Message)

	log.Println(" Send Status Request and handle stream ")
	client := api.NewStatusStreamClient(conn)
	streamer, err := client.Status(ctx, &api.Request{})
	if err != nil {
		log.Fatalf("Error when calling NewStatusStreamClient: %s", err)
	}
	msgChan := make(chan *api.StatusResponse)
	errChan := make(chan error)
	wg := sync.WaitGroup{}
	go func() {
		for errChan != nil && msgChan != nil {
			msg, err := streamer.Recv()
			if err != nil {
				errChan <- err
				return
			}
			msgChan <- msg
		}
	}()
	wg.Add(1)
	func() {
		for {
			select {
			case msg := <-msgChan:
				fmt.Println("Message received is ", msg)
			case err := <-errChan:
				fmt.Println("Error receiving ", err)
				return
			case <-ctx.Done():
				fmt.Println("context cancelled done")
				return
			}
		}
	}()
	close(msgChan)
	close(errChan)
	msgChan = nil
	errChan = nil
	wg.Done()
}
