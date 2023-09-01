package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("Long greet function was invoked")
	res := ""

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalln("Error while reading client stream: ", err)
		}
		res += fmt.Sprintf("Hello %s!\n", resp.FirstName)
	}
}
