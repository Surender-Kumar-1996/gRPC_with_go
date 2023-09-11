package main

import (
	"context"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("-------update blog was invoked----------")
	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "not Surender",
		Title:    "A new title",
		Content:  "Bekhe Bekhe kadam, behka behka hai man, cha gaya cha gaya mujhpe diwanapan",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalln("Error happened while updating: ", err)
	}

	log.Println("Blog was updated!")
}
