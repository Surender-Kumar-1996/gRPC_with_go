package main

import (
	"io"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/calculator/proto"
)

func (s SumServer) Max(stream pb.SumService_MaxServer) error {
	log.Println("Max invoked on server")
	for {
		max := 0
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error occured while receiving request from client side: ", err)
		}
		if max < int(req.Num) {
			max = int(req.Num)
			err := stream.Send(&pb.MaxResponse{Maxx: req.Num})
			if err != nil {
				log.Fatalln("Error occured while sending data to the client: ", err)
			}
		}
	}
	return nil
}
