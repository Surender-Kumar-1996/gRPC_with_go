package main

import (
	"log"
	"net"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Faild to listen on,", addr, " with error: ", err.Error())
	}
	log.Println("Listning on address:", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalln("FAiled to serve: ", err)
	}

}
