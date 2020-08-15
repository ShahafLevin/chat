package connector

import (
	"chat/framework/communicator"
	"chat/framework/connector"
	"chat/framework/structs"
)

// Connector represents a client connector, which holds the methods in
// the connector.Connector and a method to connect the Room
type Connector interface {
	connector.Connector
	ConnectToRoom(structs.RoomID) (communicator.Communicator, error)
}
