package communicator

import (
	"chat/framework/message"
)

// Communicator is an IO abstraction in our chat ecosystem
type Communicator interface {
	// Send read messages from a channel, converts it to bytes and writes it to a connection
	Send(chan message.Message) error
	// Recieve read bytes from a connection, converts it to message and sends it to a given message channel
	Recieve(chan message.Message) error
}
