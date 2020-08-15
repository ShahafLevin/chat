package connector

import (
	"bufio"
	"chat/framework/communicator"
	"chat/framework/structs"
	"fmt"
	"net"
)

// netConnector represnts a client net connector, which holds the same attributes
// as connector.NetConnector
type netConnector struct {
	conn net.Conn
}

// NewNetConnector creates new net connector
func NewNetConnector(addr string, method string) (Connector, error) {
	conn, err := net.Dial(method, addr)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial. %s", err)
	}
	return &netConnector{conn: conn}, nil
}

// ConnectToRoom connects to a room using net protocol
func (nc *netConnector) ConnectToRoom(room structs.RoomID) (communicator.Communicator, error) {
	nc.conn.Write([]byte(room + "\n"))
	repsonse, err := bufio.NewReader(nc.conn).ReadByte()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect the room: %s", err)
	}
	if repsonse == '1' {
		return nil, fmt.Errorf("Room %s does not exist", room)
	}
	return nc.Connect()
}

// Connect inits the net communicator
func (nc *netConnector) Connect() (communicator.Communicator, error) {
	return communicator.NewNetCommunicator(nc.conn, false)
}
