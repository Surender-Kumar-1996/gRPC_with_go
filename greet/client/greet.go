package main

import (
	"context"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("DoGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Client",
	})

	if err != nil {
		log.Fatalln("Could not greet: ", err.Error())
	}

	log.Println("Greeting ", res.Result)
}
