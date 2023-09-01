package main

import (
	"fmt"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "127.0.0.1:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("Dialing to address ", addr)
	if err != nil {
		log.Fatalln("Failed to connect.....", err.Error())
	}
	defer conn.Close()

	c := pb.NewSumServiceClient(conn)
	fmt.Println("Calling do sum")
	// doSum(c)
	// doPrime(c)
	// doAverage(c)
	doMax(c)
	fmt.Println("DoSum call completed")
}
