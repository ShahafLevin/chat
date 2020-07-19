package client

import (
	"bufio"
	"chat/framework/message"
	"chat/framework/user"
	"fmt"
	"io"
	"log"
	"os"
)

// RoomID represents the room ID
type RoomID string

type chatClient interface {
	Run() error
	Write()
	Read()
}

// Client - a struct for a client on the chat server
type Client struct {
	connector Connector
	room      RoomID
}

// Run is the method which runs the client
func (client *Client) Run() error {
	if err := client.connector.ConnectToRoom(client.room); err != nil {
		return err
	}
	log.Println("Connected to room: ", client.room)

	go client.Write()
	client.Read()
	return nil
}

func (client *Client) Write() error {
	msgChan := make(chan message.Message)
	userInput := bufio.NewReader(os.Stdin)
	client.connector.Send(msgChan)
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
	client.connector.Recieve(msgChan)
	for {
		fmt.Print(<-msgChan)
	}
}
