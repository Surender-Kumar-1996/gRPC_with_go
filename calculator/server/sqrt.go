package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *SumServer) Sqrt(ctx context.Context, req *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Println("Square root function called with: ", req)
	if req.Number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Received a negetive number: %d", req.Number))
	}
	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(req.Number)),
	}, nil
}
