package main

import (
	"io"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/calculator/proto"
)

func (s *SumServer) Avg(stream pb.SumService_AvgServer) error {
	log.Println("Average endpoint was called")

	res := int32(0)
	len := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Calculating average ", res, "/", len)
			average := float64(res) / float64(len)
			stream.SendAndClose(&pb.AvgResponse{
				Result: average,
			})
			return nil
		}
		if err != nil {
			log.Fatalln("Error occured while receiving client stream:", err)
		}
		res += int32(req.Num)
		len += int32(1)
	}
}
