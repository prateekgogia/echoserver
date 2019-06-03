package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/prateekgogia/echoserver/pkg/server"
)

var (
	port = flag.Int("port", 8080, "gRPC server port")
)

// main start a gRPC server
func main() {
	flag.Parse()
	serv := server.New()
	// handle interrupt and shutdown the server
	go handleInterrupt(serv)
	// start server, blocking call
	if err := serv.Run(*port); err != nil {
		log.Fatalf("failed to start gRPC server err %v", err)
	}
}

func handleInterrupt(s *server.Server) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	log.Printf("Received an interrupt, stopping services...\n")
	s.Stop()
	log.Printf("Done .... \n")
}
