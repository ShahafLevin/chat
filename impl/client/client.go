package client

import (
	"bufio"
	"chat/framework/connector"
	"chat/framework/message"
	"chat/framework/user"
	"fmt"
	"io"
	"log"
	"os"
)

type chatClient interface {
	Run() error
	Write()
	Read()
}

// Client - a struct for a client on the chat server
type Client struct {
	Connector Connector
	Room      connector.RoomID
}

// Run is the method which runs the client
func (client *Client) Run() error {
	if err := client.Connector.ConnectToRoom(client.Room); err != nil {
		return err
	}
	log.Println("Connected to room: ", client.Room)

	go client.Write()
	client.Read()
	return nil
}

func (client *Client) Write() error {
	msgChan := make(chan message.Message)
	userInput := bufio.NewReader(os.Stdin)
	client.Connector.Send(msgChan)
	for {
		userLine, err := userInput.ReadBytes(byte('\n'))
		switch err {
		case nil:
			// Todo: init a User as well
			var user user.User
			msgChan <- message.NewText(userLine, user)
		case io.EOF:
			return fmt.Errorf("No more output to send to connection")
		default:
			return fmt.Errorf("Somthing wrong happend %s", err)
		}
	}
}

func (client *Client) Read() {
	msgChan := make(chan message.Message)
	client.Connector.Recieve(msgChan)
	for {
		fmt.Print(<-msgChan)
	}
}
