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
	addr   string
	method string
}

// NewNetConnector creates new net connector
func NewNetConnector(addr string, method string) Connector {
	return &netConnector{addr: addr, method: method}
}

// ConnectToRoom connects to a room using net protocol
func (nc *netConnector) ConnectToRoom(room structs.RoomID) (communicator.Communicator, error) {
	conn, err := net.Dial(nc.method, nc.addr)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial. %s", err)
	}

	conn.Write([]byte(room + "\n"))
	repsonse, err := bufio.NewReader(conn).ReadByte()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect the room: %s", err)
	}
	if repsonse == '1' {
		return nil, fmt.Errorf("Room %s does not exist", room)
	}
	return nc.Connect(conn)
}

// Connect inits the net communicator
func (nc *netConnector) Connect(conn interface{}) (communicator.Communicator, error) {
	return communicator.NewNetCommunicator(conn.(net.Conn), false)
}
