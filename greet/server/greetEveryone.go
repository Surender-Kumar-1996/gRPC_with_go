package main

import (
	"io"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/greet/proto"
)

func (s *Server) GreetEveryOne(stream pb.GreetService_GreetEveryOneServer) error {
	log.Println("Everyone was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalln("Error while reading client stream: ", err)
		}
		res := "Hello" + req.FirstName + "!\n"
		err = stream.Send(&pb.GreetResponse{Result: res})
		if err != nil {
			log.Println("Unable to send response to the client: ", err)
		}
	}
}
