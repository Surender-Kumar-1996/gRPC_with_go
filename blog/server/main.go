package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var collection *mongo.Collection

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	//making the Database connection
	t := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalln("err")
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection = client.Database("blogdb").Collection("blog")

	//Creating a listner on addr
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen on: ", err)
	}

	log.Println("Listning on ", addr)

	//Implemention ssl for the server
	ssl := true
	opts := []grpc.ServerOption{}

	if ssl {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"

		cert, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalln("Failed loading certificate: ", err.Error())
		}
		opts = append(opts, grpc.Creds(cert))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterBlogServiceServer(s, &Server{})

	//Server the server on addr
	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve: ", err)
	}
}
