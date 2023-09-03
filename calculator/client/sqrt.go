package main

import (
	"context"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.SumServiceClient, n int32) {
	log.Println("doSqrt called....")

	resp, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Println("Error msg from server", e.Message(), " wiht error code", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalln("A non gRPC err: ", err.Error())
		}
	}

	log.Println("SquareRoot of ", n, " is: ", resp.Result)
}
