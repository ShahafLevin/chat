package connector

import (
	"bufio"
	"chat/framework/communicator"
	"chat/framework/structs"
	"chat/impl/server/room"
	"fmt"
	"log"
	"net"
)

type netConnector struct {
	port     int
	method   string
	listener net.Listener
}

// NewNetConnector creates new net connector
func NewNetConnector(port int, method string) Connector {
	return &netConnector{
		port:   port,
		method: method,
	}
}

func (nc *netConnector) Serve(rooms room.Rooms) error {
	if err := nc.initListener(); err != nil {
		return err
	}
	defer nc.listener.Close()

	for {
		if err := nc.handleUser(rooms); err != nil {
			log.Print(err)
		}
	}
}

func (nc *netConnector) Connect(conn interface{}) (communicator.Communicator, error) {
	return communicator.NewNetCommunicator(conn.(net.Conn), true)
}

func (nc *netConnector) AddToRoom(room room.Room, comm communicator.Communicator) error {
	return room.AddComm(comm)
}

func (nc *netConnector) initListener() (err error) {
	nc.listener, err = net.Listen(nc.method, fmt.Sprintf(":%d", nc.port))
	if err != nil {
		return err
	}
	return
}

func (nc *netConnector) handleUser(rooms room.Rooms) error {
	conn, err := nc.listener.Accept()
	if err != nil {
		return err
	}
	// Todo: Close connections

	roomAsBytes, err := bufio.NewReader(conn).ReadBytes(byte('\n'))
	if err != nil {
		return err
	}

	roomID := structs.RoomID(roomAsBytes[:len(roomAsBytes)-1])
	room, ok := rooms[roomID]
	if !ok {
		conn.Write([]byte{'1'})
		return nil
	}
	conn.Write([]byte{'2'})

	comm, err := nc.Connect(conn)
	if err != nil {
		return err
	}

	if err = nc.AddToRoom(room, comm); err != nil {
		return err
	}

	return nil

}
