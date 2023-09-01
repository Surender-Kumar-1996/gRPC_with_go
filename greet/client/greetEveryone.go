package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("DoGreetEveryone was invoked")

	stream, err := c.GreetEveryOne(context.Background())
	if err != nil {
		log.Fatalln("Error while creating stream: ", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Surender"},
		{FirstName: "Pankaj"},
		{FirstName: "Pritam"},
		{FirstName: "Bhavesh"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Println("Send Request: ", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("Error while receiving response from server: ", err)
			}
			log.Println("Received: ", res.Result)
		}
		close(waitc)
	}()

	<-waitc

}
