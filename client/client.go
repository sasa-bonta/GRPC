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
	"time"
)

func main() {
	fmt.Println("Starting client")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(common.HostPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	c := client.NewChatServiceClient(conn)

	defer func(conn *grpc.ClientConn) {
		message := client.Message{
			Action: common.UNSUBSCRIBE,
		}
		_, err := c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		err = conn.Close()
		if err != nil {
			fmt.Println("Cannot close connection")
			os.Exit(1)
		}
	}(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter topic: ")
	topic, _ := reader.ReadString('\n')

	message := client.Message{
		Action: common.SUBSCRIBE,
		Topic:  strings.ToLower(strings.TrimSpace(topic)),
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: {action : \"%s\", topic : \"%s\", body : \"%s\"}", response.Action, response.Topic, response.Body)

	message.Action = common.TEST

	for {
		time.Sleep(3 * time.Second)
		response, err := c.SayHello(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}

		if response.Body != "" {
			log.Printf("New message: %s", response.Body)
		}
	}
}
