package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/blog/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("List Blog service was called..")
	stream, err := c.ListBlog(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalln("Error while calling ListBlog:", err.Error())
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Something happned ", err.Error())
		}
		log.Println(res)
	}
}
