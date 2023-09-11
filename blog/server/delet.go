package main

import (
	"context"
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/blog/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*empty.Empty, error) {
	log.Println("DeleteBlog was invoked with", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot parse ID: "+err.Error())
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot delete object in MongoDB: ", err.Error())
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.Internal, "Blog was not found")
	}
	return &emptypb.Empty{}, nil
}
