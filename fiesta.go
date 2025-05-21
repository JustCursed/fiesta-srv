package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"void-studio.net/fiesta/config"
	"void-studio.net/fiesta/pb"
	"void-studio.net/fiesta/services"
)

func main() {
	listener, err := net.Listen("tcp4", config.Config.General.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterCollectorServer(server, &services.Collector{})

	fmt.Println("Server started")

	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
