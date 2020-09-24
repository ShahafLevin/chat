// Package connector purpus is to create communicators, in which the client and server
// will use in order to communicate eachother. The main idea is to create an abstraction to the
// handshake between the client and the server.
package connector

import (
	"chat/framework/communicator"
)

// Connector represents a connector in the chat
type Connector interface {
	Connect(interface{}) (communicator.Communicator, error)
}
