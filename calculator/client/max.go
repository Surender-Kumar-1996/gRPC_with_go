package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/calculator/proto"
)

func doMax(c pb.SumServiceClient) {
	log.Println("DoMax service was invoked")

	arr := []int32{1, 5, 3, 6, 2, 20}

	client, err := c.Max(context.Background())
	if err != nil {
		log.Fatalln("Error while calling Max from server side", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, x := range arr {
			log.Println("Sending ", x, "to check max")
			client.Send(&pb.MaxRequest{Num: x})
			time.Sleep(1 * time.Second)
		}
		err := client.CloseSend()
		if err != nil {
			log.Fatalln("Error while closing the close send:", err)
		}
	}()

	go func() {
		for {
			resp, err := client.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("Error occured while receiving stream fomr server", err)
			}
			log.Println("Max is: ", resp.Maxx)
		}
		close(waitc)
	}()

	<-waitc

}
