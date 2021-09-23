package main

import (
	"bufio"
	"fmt"
	"github.com/tutorialedge/go-grpc-tutorial/broker/cmd/storage/client"
	"github.com/tutorialedge/go-grpc-tutorial/common"
	"log"
	"os"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting publisher")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(common.HostPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := client.NewChatServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter topic: ")
		topic, _ := reader.ReadString('\n')

		fmt.Print("Enter message: ")
		body, _ := reader.ReadString('\n')

		message := client.Message{
			Action: common.PUBLISH,
			Topic:  strings.ToLower(topic),
			Body:   strings.ToLower(body),
		}

		response, err := c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}

		log.Printf("Response from Server: {action : \"%s\", topic : \"%s\", body : \"%s\"}", response.Action, response.Topic, response.Body)
	}
}
