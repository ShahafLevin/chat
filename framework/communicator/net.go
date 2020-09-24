package communicator

import (
	"bufio"
	"chat/framework/cryptochat"
	"chat/framework/message"
	"chat/framework/structs"
	"fmt"
	"io"
	"net"
)

// netCommunicator represents a connection via net.Conn
type netCommunicator struct {
	conn net.Conn
	key  cryptochat.Key
}

// NewNetCommunicator cretaes new netCommunicator
func NewNetCommunicator(conn net.Conn, isServer bool) (Communicator, error) {
	secret, err := cryptochat.EstablishSecret(conn, isServer)
	if err != nil {
		return nil, fmt.Errorf("Failed to establish secret. %s", err)
	}

	return &netCommunicator{
		conn: conn,
		key:  secret,
	}, nil
}

// Send reads from an incoming channel of messages and send it through the connector
func (nc *netCommunicator) Send(stdin chan message.Message) error {
	for {
		msg := <-stdin
		// EncryptMessage mutates the input!
		content := append(cryptochat.EncryptMessage(nc.key, msg.Marshal()), byte('\n'))
		if _, err := nc.conn.Write(content); err != nil {
			return err
		}
	}
}

// Recieve reads messages from the connector connection and sends it to a given channel
func (nc *netCommunicator) Recieve(stdout chan message.Message) error {
	stream := bufio.NewReader(nc.conn)
	for {
		content, err := stream.ReadBytes(byte('\n'))
		switch err {
		case nil:
			decrypted := cryptochat.DecryptMessage(nc.key, content[:len(content)-1])
			msg := message.NewText(decrypted, structs.User{})
			stdout <- msg
		case io.EOF:
			return fmt.Errorf("No more input to read from connection")
		default:
			return fmt.Errorf("Somthing wrong happend. %s ", err)
		}
	}
}
