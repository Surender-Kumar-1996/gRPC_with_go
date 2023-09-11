package main

import (
	"context"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/blog/proto"
)

func createBlogs(c pb.BlogServiceClient) string {
	log.Println("Create blog was invoked...")

	blog := &pb.Blog{
		AuthorId: "Surender",
		Title:    "My first Blog",
		Content:  "Chikni chamali chup ke akali pauua chadha ke aayi",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatal("Unexpected err: " + err.Error())
	}

	log.Println("Blog created successfully with id: " + res.Id)

	return string(res.Id)

}
