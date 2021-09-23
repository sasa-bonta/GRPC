package chat

import (
	"github.com/tutorialedge/go-grpc-tutorial/common"
	"github.com/tutorialedge/go-grpc-tutorial/storage"
	"golang.org/x/net/context"
	"log"
)

var ms = &storage.MessagesList{}

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Server received message: {action : \"%s\", topic : \"%s\", body : \"%s\"}", message.Action, message.Topic, message.Body)
	messageToAdd := storage.Message{
		Action: message.Action,
		Topic: message.Topic,
		Body: message.Body,
	}
	ms.AddMessage(&messageToAdd)
	return &Message{Action: common.CONFIRM, Topic: common.SUCCESS}, nil
}
