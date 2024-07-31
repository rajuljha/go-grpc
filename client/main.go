package main

import (
	"log"

	pb "github.com/rajuljha/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = "8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to the client %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// names := &pb.NameList{
	// 	Names: []string{"Rajul", "Alice", "Bob"},
	// }

	// callSayHello(client)
	
}
