package main

import (
	"fmt"
	"github.com/tutorialedge/go-grpc-tutorial/chat"
	"github.com/tutorialedge/go-grpc-tutorial/common"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting server")
	lis, err := net.Listen(common.ConnectionType, common.HostPort)
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}
