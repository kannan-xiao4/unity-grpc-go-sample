package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "sample/pb"
	controller "sample/controller"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", "10000"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	chatServer := &controller.ChatController{}
	pb.RegisterRouteGuideServer(grpcServer, chatServer)
	grpcServer.Serve(lis)
}