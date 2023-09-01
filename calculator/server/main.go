package main

import (
	"log"
	"net"

	pb "github.com/Surender-Kumar-1996/gRPC_with_go/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "127.0.0.1:50051"

type SumServer struct {
	pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Faild to listen on,", addr, " with error: ", err.Error())
	}
	log.Println("Listning on address:", addr)

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &SumServer{})

	if err = s.Serve(lis); err != nil {
		log.Fatalln("FAiled to serve: ", err)
	}

}
