package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/calculator/proto"
)

func doSum(c pb.SumServiceClient) {
	fmt.Println("************Hi caught up here****************")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 1,
	})

	if err != nil {
		log.Fatalln("could not Sum: ", err.Error())
	}
	fmt.Println("Here i am")

	log.Println("Calculation result is: ", res.Result)
}
