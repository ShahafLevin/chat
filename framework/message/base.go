package message

import "chat/framework/user"

// Message represents a message in the chat server
type Message interface {
	Marshal() []byte
	UnMarshal([]byte)
}

// BaseMessage used to represent a base message
type BaseMessage struct {
	Content []byte
	Sender  user.User
	Type    string
}
