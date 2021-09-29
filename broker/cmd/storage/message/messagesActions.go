package message

import (
	"fmt"
)

func (ml *MessagesList) AddMessage(message *Message) *MessagesList {
	ml.Messages = append(ml.Messages, *message)
	fmt.Println("Added new message")
	return ml
}

func (ml *MessagesList) NextMessage() *Message {
	fmt.Println("Shifting message")
	if !ml.IsEmpty() {
		message := ml.Messages[0]
		ml.Messages = ml.Messages[1:]
		return &message
	}
	return nil
}

func (ml *MessagesList) IsEmpty() bool {
	if len(ml.Messages) == 0 {
		return true
	}
	return false
}
