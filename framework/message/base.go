package message

import "chat/framework/structs"

// Message represents a message in the chat server
type Message interface {
	Marshal() []byte
	UnMarshal([]byte)
	User() structs.User
}

// BaseMessage used to represent a base message
type BaseMessage struct {
	Content []byte
	Sender  structs.User
	Type    string
}
