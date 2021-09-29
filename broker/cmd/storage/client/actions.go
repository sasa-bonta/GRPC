package client

import (
	"fmt"
	"github.com/tutorialedge/go-grpc-tutorial/common"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
)

func (cl *ClientsList) Publish(topic, message string) {
	subscriptions := cl.GetSubscriptions(topic, nil)

	for _, sub := range subscriptions {
		fmt.Printf("Sending to client: %s message is %s \n", sub.Address, message)
		messageObj := Message{Action: common.PUBLISH, Topic: topic, Body: message}
		sendMessage(sub.Address, &messageObj)
	}
}

func (cl *ClientsList) GetSubscriptions(topic string, client *Client) []Client {

	var subscriptionList []Client

	for _, subscription := range cl.Clients {

		if client != nil {
			if subscription.Address != "" && subscription.Topic == topic {
				subscriptionList = append(subscriptionList, subscription)
			}
		} else {
			if subscription.Topic == topic {
				subscriptionList = append(subscriptionList, subscription)
			}
		}
	}

	return subscriptionList
}

func sendMessage(address string, messageObj *Message) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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

	c := NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), messageObj)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response fromClient: {action : \"%s\", topic : \"%s\", body : \"%s\"}", response.Action, response.Topic, response.Body)
}
