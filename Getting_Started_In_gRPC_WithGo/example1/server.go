package main

import (
	"log"
	"net"

	"example1/pb"
	"example1/server"

	"google.golang.org/grpc"
)

func main() {
	// Created a listener
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create new server and register our services
	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, &server.Calculator{})

	// Start server with listener
	log.Println("Server starting...")
	log.Fatal(grpcServer.Serve(lis))
}
