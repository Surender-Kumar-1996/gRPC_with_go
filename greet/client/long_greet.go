package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "vickey"},
		{FirstName: "Surender"},
		{FirstName: "Sunny"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalln("Error while calling long greet: ", err)
	}

	for _, req := range reqs {
		log.Println("Sending req: ", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error while receiving response from Longgreet:", err)
	}

	log.Println("LongGreet Result: ", res.Result)
}
