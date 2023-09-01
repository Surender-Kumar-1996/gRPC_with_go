package main

import (
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/calculator/proto"
)

func (s *SumServer) Prime(in *pb.PrimeRequest, stream pb.SumService_PrimeServer) error {
	log.Println("Prime number function was invoked with: ", in)

	var k int32 = 2
	n := in.Numb
	for n > 1 {
		if n%k == 0 {
			res := &pb.SumResponse{
				Result: k,
			}
			stream.Send(res)
			n = n / k
		} else {
			k = k + 1
		}

	}
	return nil
}
