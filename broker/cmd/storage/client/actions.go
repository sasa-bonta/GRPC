package client

import (
	"fmt"
	"github.com/tutorialedge/go-grpc-tutorial/chat"
	"github.com/tutorialedge/go-grpc-tutorial/common"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func (cl *ClientsList) Publish(topic, message string) {
	subscriptions := cl.GetSubscriptions(topic, nil)

	for _, sub := range subscriptions {
		fmt.Printf("Sending to client id %s message is %s \n", sub.Address, message)
		messageObj := chat.Message{Action: common.PUBLISH, Topic: topic, Body: message}
		conn, err := grpc.Dial(common.HostPort, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("could not connect to client: %s", err)
		}
		c := chat.NewChatServiceClient(conn)
		c.SayHello(context.Background(), &messageObj)
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
