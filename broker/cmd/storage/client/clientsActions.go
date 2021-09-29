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

func (cl *ClientsList) RemoveClient(address string) {
	fmt.Println("Removing client")
	// bring element to remove at the end if its not there yet
	i := cl.index(address)
	if i != len(cl.Clients)-1 {
		cl.Clients[i] = cl.Clients[len(cl.Clients)-1]
	}

	// drop the last element
	cl.Clients = cl.Clients[:len(cl.Clients)-1]
}

func (cl *ClientsList) index(address string) int {
	for i := range cl.Clients {
		if cl.Clients[i].Address == address {
			return i
		}
	}
	return -1
}
