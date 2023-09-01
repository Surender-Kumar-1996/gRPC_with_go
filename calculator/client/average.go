package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/calculator/proto"
)

func doAverage(c pb.SumServiceClient) {
	log.Println("Average function was invoked")

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalln("Error while calling Average:", err)
	}

	for i := 1; i <= 5; i++ {
		err = stream.Send(&pb.StreamAverageRequest{
			Num: int32(i),
		})
		if err != nil {
			log.Fatalln("Unable to send stream of data to the Average endpoint: ", err)
		}
		log.Println("Sending stream request with ", i)
		time.Sleep(1 * time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Unable to close and receive the stream avg request :", err)
	}

	log.Println("Average of 1,2,3,4,5 is ", resp.Result)
}
