package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("Greet with deadline was invoked with: ", in)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Client canceled the request")
			return nil, status.Error(codes.Canceled, "The Client cancelled the request")
		}

		time.Sleep(1 * time.Second)
	}
	return &pb.GreetResponse{Result: "Hello " + in.FirstName + "!!"}, nil
}
