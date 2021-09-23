package main

import (
	"fmt"
	"github.com/tutorialedge/go-grpc-tutorial/broker/cmd/storage/client"
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

	s := client.Server{}

	grpcServer := grpc.NewServer()

	client.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}
