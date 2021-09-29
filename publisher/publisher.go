package main

import (
	"bufio"
	"fmt"
	"github.com/tutorialedge/go-grpc-tutorial/broker/cmd/storage/client"
	"github.com/tutorialedge/go-grpc-tutorial/common"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting publisher")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(common.HostPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Cannot close connection")
			os.Exit(1)
		}
	}(conn)

	c := client.NewChatServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter topic: ")
		topic, _ := reader.ReadString('\n')

		fmt.Print("Enter message: ")
		body, _ := reader.ReadString('\n')

		message := client.Message{
			Action: common.PUBLISH,
			Topic:  strings.ToLower(strings.TrimSpace(topic)),
			Body:   strings.ToLower(strings.TrimSpace(body)),
		}

		response, err := c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}

		log.Printf("Response from Server: {action : \"%s\", topic : \"%s\", body : \"%s\"}", response.Action, response.Topic, response.Body)
	}

}
