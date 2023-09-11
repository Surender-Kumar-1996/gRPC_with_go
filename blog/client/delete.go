package main

import (
	"context"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("Delete Blog was called...")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalln("Error while deleting the blog:", err.Error())
	}

	log.Println("Blog was Deleted...")
}
