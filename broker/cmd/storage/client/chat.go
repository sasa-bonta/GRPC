package client

import (
	"errors"
	messageDir "github.com/tutorialedge/go-grpc-tutorial/broker/cmd/storage/message"
	"github.com/tutorialedge/go-grpc-tutorial/common"
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
	"log"
)

var ms = &messageDir.MessagesList{}
var cs = &ClientsList{}

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {

	if message.Topic == "" {
		return &Message{}, errors.New("empty name")
	}

	log.Printf("Server received message: {action : \"%s\", topic : \"%s\", body : \"%s\"}", message.Action, message.Topic, message.Body)

	receivedMessage := messageDir.Message{
		Action: message.Action,
		Topic:  message.Topic,
		Body:   message.Body,
	}

	p, _ := peer.FromContext(ctx)
	address := p.Addr.String()

	switch message.Action {

	case common.PUBLISH:
		ms.AddMessage(&receivedMessage)
		break
	case common.SUBSCRIBE:
		cs.AddSubscription(message.Topic, address)
		break
	case common.UNSUBSCRIBE:
		cs.RemoveClient(address)
		break
	case common.TEST:
		for !ms.IsEmpty() {
			publishMessage := ms.NextMessage()
			if message.Topic == publishMessage.Topic {
				log.Printf("Sending: {action : \"%s\", topic : \"%s\", body : \"%s\"}", publishMessage.Action, publishMessage.Topic, publishMessage.Body)
				return &Message{Action: common.PUBLISH, Topic: publishMessage.Topic, Body: publishMessage.Body}, nil
			}
		}

		break
	default:
		break
	}
	return &Message{Action: common.CONFIRM, Topic: common.SUCCESS}, nil
}
