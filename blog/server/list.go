package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/blog/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListBlog(empt *empty.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Println("ListBlog was invoked...")

	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unkonwn internal error: %v", err.Error()),
		)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)
		if err != nil {
			status.Errorf(codes.Internal, "Error while decoding data from MongoDB: "+err.Error())
		}

		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		status.Errorf(codes.Internal, "Unkonwn internal error: "+err.Error())
	}

	return nil
}
