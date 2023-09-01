package main

import (
	"fmt"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/greet/proto"
)

func (s Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Println("GreetManyTimes function was invoked with ", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
