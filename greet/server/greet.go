package main

import (
	"context"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("Greet function was invoked with ", in)
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
