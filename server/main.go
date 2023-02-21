package main

import (
	"net"

	"google.golang.org/grpc"

	pb "github.com/prashantkumardagur/grpc-go/proto"
)

//==============================================================================

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Server struct {
	pb.UnimplementedGreetServiceServer
}

//==============================================================================

func main() {
	// Create a TCP listener
	tpcServer, err := net.Listen("tcp", ":8080")
	HandleErr(err)

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Register the service with the gRPC server
	pb.RegisterGreetServiceServer(grpcServer, &Server{})

	// Register the gRPC server
	grpcErr := grpcServer.Serve(tpcServer)
	HandleErr(grpcErr)
}
