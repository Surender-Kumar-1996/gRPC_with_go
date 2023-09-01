package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/calculator/proto"
)

func doPrime(c pb.SumServiceClient) {
	log.Println("doPrime was invoked")

	res := &pb.PrimeRequest{
		Numb: 120,
	}

	stream, err := c.Prime(context.Background(), res)
	if err != nil {
		log.Fatalln("Error while calling prime:", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error while receiving stream: ", err)
		}
		log.Print("Prime factor of ", res.Numb, "is ", resp.Result)
	}
}
