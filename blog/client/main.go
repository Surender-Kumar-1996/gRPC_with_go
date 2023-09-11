package main

import (
	"log"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalln("Error while loading CA trust certificate:", err.Error())
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalln("Failed to connect:" + err.Error())
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlogs(c)
	readBlog(c, id) //valid
	// readBlog(c, "123erfdsd543wd")

	updateBlog(c, id)

	listBlog(c)

	deleteBlog(c, id)

}
