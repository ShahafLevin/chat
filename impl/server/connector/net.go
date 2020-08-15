package connector

import (
	"chat/framework/communicator"
	"chat/impl/server/room"
	"net"
)

type netConnector struct {
	listener net.Listener
	conn     chan net.Conn
	rooms    room.Rooms
}

// NewNetConnector creates new net connector
func NewNetConnector(port int, method string, rooms room.Rooms) (Connector, error) {
	return nil, nil
}

func (nc *netConnector) Listen() (chan communicator.Communicator, error) {
	return nil, nil
}

func (nc *netConnector) Connect() (communicator.Communicator, error) {
	return nil, nil
}

func (nc *netConnector) AddToRoom(room room.Room, comm communicator.Communicator) error {
	return nil
}
