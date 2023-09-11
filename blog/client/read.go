package main

import (
	"context"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("-------Read Blog was Invoked-------")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Println("Error happned while reading: ", err.Error())
		return nil
	}

	log.Println("Blog was read: ", res)
	return res
}
