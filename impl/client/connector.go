package client

import (
	"bufio"
	"chat/framework/cryptochat"
	"chat/framework/message"
	"fmt"
	"io"
	"net"
)

// Connector represents a connector to a remote chat server
type Connector interface {
	ConnectToRoom(RoomID) error
	Send(chan message.Message) error
	Recieve(chan message.Message) error
}

// NetConnector represents a connection via net.Conn
type NetConnector struct {
	Address string
	Conn    net.Conn
	Key     cryptochat.Key
}

// NewNetConnector cretaes new NetConnector
func NewNetConnector(addr string, method string) (*NetConnector, error) {
	conn, err := net.Dial(method, addr)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial. %s", err)
	}

	secret, err := establishSecret(conn)
	if err != nil {
		return nil, fmt.Errorf("Failed to establish secret. %s", err)
	}

	return &NetConnector{
		Address: addr,
		Conn:    conn,
		Key:     secret,
	}, nil
}

// ConnectToRoom connects to a room using net protocl
func (connector *NetConnector) ConnectToRoom(room RoomID) error {
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

// Send reads from an incoming channel of messages and send it through the connector
func (connector *NetConnector) Send(stdin chan message.Message) error {
	for {
		msg := <-stdin
		content := append(cryptochat.EncryptMessage(connector.Key, msg.Marshal()), byte('\n'))
		if _, err := connector.Conn.Write(content); err != nil {
			return err
		}
	}
}

// Recieve reads messages from the connector connection and sends it to a given channel
func (connector *NetConnector) Recieve(stdout chan message.Message) error {
	stream := bufio.NewReader(connector.Conn)
	for {
		content, err := stream.ReadBytes(byte('\n'))
		switch err {
		case nil:
			var msg message.Message
			msg.UnMarshal(cryptochat.DecryptMessage(connector.Key, content[:len(content)-1]))
			stdout <- msg
		case io.EOF:
			return fmt.Errorf("No more input to read from connection")
		default:
			return fmt.Errorf("Somthing wrong happend. %s ", err)
		}
	}
}

// establishSecret establish secret with the user
// TODO: add error in case of somthing went wrong
func establishSecret(conn net.Conn) (secert []byte, err error) {
	key := cryptochat.GenerateKey()
	serverKey := cryptochat.ReadKey(conn, (*key))
	cryptochat.WriteKey(conn, (*key))
	return key.KeyExchange.ComputeSecret(key.Private, serverKey), nil
}
