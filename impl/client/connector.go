package client

import (
	"bufio"
	"chat/framework/connector"
	"fmt"
)

// Connector represents a client connector
type Connector interface {
	connector.Connector
	ConnectToRoom(connector.RoomID) error
}

// NetConnector represnts a client net connector
type NetConnector struct {
	connector.NetConnector
}

// NewNetConnector creates new net connector
func NewNetConnector(addr string, method string, server bool) (*NetConnector, error) {
	conn, err := connector.NewNetConnector(addr, method, server)
	return &NetConnector{(*conn)}, err
}

// ConnectToRoom connects to a room using net protocl
func (connector *NetConnector) ConnectToRoom(room connector.RoomID) error {
	connector.Conn.Write([]byte(room + "\n"))
	repsonse, err := bufio.NewReader(connector.Conn).ReadByte()
	if err != nil {
		return fmt.Errorf("Failed to connect the room: %s", err)
	}
	if repsonse == '1' {
		return fmt.Errorf("Room %s does not exist", room)
	}
	return nil
}
