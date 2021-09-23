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
	"time"
)

func main() {
	fmt.Println("Starting client")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(common.HostPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := client.NewChatServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter topic: ")
	topic, _ := reader.ReadString('\n')

	message := client.Message{
		Action: common.SUBSCRIBE,
		Topic:  topic,
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: {action : \"%s\", topic : \"%s\", body : \"%s\"}", response.Action, response.Topic, response.Body)

	time.Sleep(10 * time.Second)

	log.Println("Connection stopped")
}
