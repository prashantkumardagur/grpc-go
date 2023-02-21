package main

import (
	"context"
	"io"

	pb "github.com/prashantkumardagur/grpc-go/proto"
)

//==============================================================================

func (s *Server) UnaryGreet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Message: "Hello " + req.Name,
	}, nil
}

//==============================================================================

func (s *Server) ServerStreamingGreet(req *pb.GreetrRequestList, stream pb.GreetService_ServerStreamingGreetServer) error {
	for _, name := range req.Names {
		res := &pb.GreetResponse{
			Message: "Hello " + name,
		}
		err := stream.Send(res)
		HandleErr(err)
	}
	return nil
}

//==============================================================================

func (s *Server) ClientStreamingGreet(stream pb.GreetService_ClientStreamingGreetServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponseList{Messages: messages})
		}
		HandleErr(err)
		messages = append(messages, "Hello "+req.Name)
	}
}

//==============================================================================

func (s *Server) BiDiStreamingGreet(stream pb.GreetService_BiDiStreamingGreetServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		HandleErr(err)
		res := &pb.GreetResponse{
			Message: "Hello " + req.Name,
		}
		sendErr := stream.Send(res)
		HandleErr(sendErr)
	}
	return nil
}
