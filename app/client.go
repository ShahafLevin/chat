package app

import (
	"bufio"
	"chat/framework/communicator"
	"chat/framework/message"
	"chat/framework/structs"
	"chat/impl/client/connector"
	"fmt"
	"io"
	"os"
)

// Client represents a client in the chat server
type Client interface {
	Run() error
	Write() error
	Read()
}

// client is a struct for a client on the chat server
type client struct {
	communicator communicator.Communicator
	user         structs.User
}

// NewClient returns a new client instance
func NewClient(connector connector.Connector, room structs.RoomID, user structs.User) (Client, error) {
	communicator, err := connector.ConnectToRoom(room)
	if err != nil {
		return nil, err
	}
	return &client{
		communicator: communicator,
		user:         user,
	}, nil
}

// Run is the method which runs the client
func (c *client) Run() error {
	go c.Write()
	c.Read()
	return nil
}

func (c *client) Write() error {
	msgChan := make(chan message.Message)
	userInput := bufio.NewReader(os.Stdin)
	go c.communicator.Send(msgChan)
	for {
		userLine, err := userInput.ReadBytes(byte('\n'))
		switch err {
		case nil:
			msgChan <- message.NewText(userLine, c.user)
		case io.EOF:
			return fmt.Errorf("No more output to send to connection")
		default:
			return fmt.Errorf("Somthing wrong happend %s", err)
		}
	}
}

func (c *client) Read() {
	msgChan := make(chan message.Message)
	go c.communicator.Recieve(msgChan)
	for {
		msg := <-msgChan
		fmt.Printf("%s\n", msg.Marshal())
	}
}
