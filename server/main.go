package main

import (
	"log"
	"net"

	pb "github.com/rajuljha/go-grpc/proto"

	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreetServiceServer
}

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to connect to server at %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
}
