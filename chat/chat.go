package chat

import (
	"github.com/tutorialedge/go-grpc-tutorial/broker/cmd/storage/client"
	messageDir "github.com/tutorialedge/go-grpc-tutorial/broker/cmd/storage/message"
	"github.com/tutorialedge/go-grpc-tutorial/common"
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
	"log"
)

var ms = &messageDir.MessagesList{}
var cs = &client.ClientsList{}

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Server received message: {action : \"%s\", topic : \"%s\", body : \"%s\"}", message.Action, message.Topic, message.Body)
	messageToAdd := messageDir.Message{
		Action: message.Action,
		Topic:  message.Topic,
		Body:   message.Body,
	}
	p, _ := peer.FromContext(ctx)
	address := p.Addr.String()

	switch message.Action {

	case common.PUBLISH:
		ms.AddMessage(&messageToAdd)

		break
	case common.SUBSCRIBE:
		cs.AddSubscription(message.Topic, address)
		break
	case common.CONFIRM:
		if message.Body == common.ERROR {
			log.Println("Client didn't receive message")
		}
		break
	default:
		break
	}
	return &Message{Action: common.CONFIRM, Topic: common.SUCCESS}, nil
}
