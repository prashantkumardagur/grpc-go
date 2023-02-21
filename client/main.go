package main

import (
	"log"

	pb "github.com/prashantkumardagur/grpc-go/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//==============================================================================

func HandleErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

//==============================================================================

func main() {
	// Create a gRPC client
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	HandleErr(err)
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewGreetServiceClient(conn)

	// UnaryGreet(client)
	// ServerStreamingGreet(client)
	// ClientStreamingGreet(client)
	BiDiStreamingGreet(client)
}
