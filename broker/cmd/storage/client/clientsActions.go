package client

import (
	"fmt"
)

func (cl *ClientsList) AddSubscription(topic string, address string) *ClientsList {
	client := Client{Topic: topic, Address: address}
	cl.Clients = append(cl.Clients, client)
	fmt.Printf("Added client %s to topic %s\n", address, topic)
	return cl
}
