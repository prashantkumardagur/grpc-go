package main

import (
	"context"
	"io"
	"log"

	pb "github.com/prashantkumardagur/grpc-go/proto"
)

//==============================================================================

func UnaryGreet(client pb.GreetServiceClient) {
	// Create a request
	req := &pb.GreetRequest{
		Name: "Prashant",
	}

	// Call the gRPC server
	res, err := client.UnaryGreet(context.Background(), req)
	HandleErr(err)

	// Print the response
	log.Printf("Response: %v", res.Message)
}

//==============================================================================

func ServerStreamingGreet(client pb.GreetServiceClient) {
	// Create a request
	req := &pb.GreetrRequestList{
		Names: []string{"Prashant", "Prajjawal", "Pankaj"},
	}

	// Call the gRPC server
	stream, err := client.ServerStreamingGreet(context.Background(), req)
	HandleErr(err)

	// Print the response
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		HandleErr(err)
		log.Printf("Response: %v", msg.Message)
	}
}

//==============================================================================

func ClientStreamingGreet(client pb.GreetServiceClient) {
	// Create a request stream
	stream, streamErr := client.ClientStreamingGreet(context.Background())
	HandleErr(streamErr)

	// Send the requests
	for _, name := range []string{"Prashant", "Prajjawal", "Pankaj"} {
		req := &pb.GreetRequest{
			Name: name,
		}
		err := stream.Send(req)
		HandleErr(err)
	}

	// Close the stream and receive the response
	res, err := stream.CloseAndRecv()
	HandleErr(err)
	log.Printf("Response: %v", res.Messages)
}

//==============================================================================

func BiDiStreamingGreet(client pb.GreetServiceClient) {
	// Create a bi-directional stream
	stream, err := client.BiDiStreamingGreet(context.Background())
	HandleErr(err)

	// Create a channel to wait for the go routine to finish
	waitc := make(chan struct{})

	// goroutine to receive the response
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			HandleErr(err)
			log.Printf("Response: %v", res.Message)
		}
		// Close the channel to indicate that the goroutine is finished
		close(waitc)
	}()

	for _, name := range []string{"Prashant", "Prajjawal", "Pankaj"} {
		req := &pb.GreetRequest{
			Name: name,
		}
		err := stream.Send(req)
		HandleErr(err)
	}

	stream.CloseSend()
	<-waitc

}
