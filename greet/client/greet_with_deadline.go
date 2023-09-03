package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeOut time.Duration) {
	log.Println("deGreetWithDeadline was invoked....")
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	resp, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{FirstName: "Surender"})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Fatalln("Deadline Exceeded!!")
			}
			log.Fatalln("Unexpected gRPC error:", err)
		}
		log.Fatalln("Non gRPC err: ", err)
	}
	log.Println("GreetWithDeadline result: ", resp.Result)
}
