package connector

import (
	"chat/framework/communicator"
	"chat/framework/connector"
	"chat/impl/server/room"
)

// Connector represents a server connector, which holds the methods in
// the connector.Connector and a method to connect the Room
type Connector interface {
	connector.Connector
	// Listen
	Listen() error
	// AdddToRoom
	AddToRoom(room.Room, communicator.Communicator) error
}
