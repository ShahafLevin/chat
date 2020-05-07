package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// RoomID represents the room ID
type RoomID string

type chatClient interface {
	Run()
	connectToRoom()
	send()
	recieve()
}

// Client - a struct for a client on the chat server
type Client struct {
	address string
	room    RoomID
	method  string
	conn    net.Conn
}

// Run is the method which runs the client
func (client *Client) Run() {
	log.Printf("Connecting to room %s at %s \n", client.room, client.address)
	conn, err := net.Dial(client.method, client.address)

	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}

	client.conn = conn
	defer client.conn.Close()

	client.connectToRoom()

	go client.send()
	client.recieve()
}

// connectToRoom connects to the server with the given room number
func (client *Client) connectToRoom() {
	client.conn.Write([]byte(client.room + "\n"))
	repsonse, err := bufio.NewReader(client.conn).ReadByte()

	if err != nil {
		log.Fatal(err)
	}

	if repsonse == '1' {
		log.Fatal("Room does not exist")
	}

	log.Println("Connected!")
}

// send reads the user input and sends it to the connection
func (client *Client) send() {
	userInput := bufio.NewReader(os.Stdin)
	for {
		userLine, err := userInput.ReadBytes(byte('\n'))
		switch err {
		case nil:
			client.conn.Write(userLine)
		case io.EOF:
			log.Fatal("No more output to send to connection")
		default:
			log.Fatal("Somthing wrong happend ", err)
		}
	}
}

// recieve recieves from the server a message and prints it
func (client *Client) recieve() {
	response := bufio.NewReader(client.conn)
	for {
		serverLine, err := response.ReadBytes(byte('\n'))
		switch err {
		case nil:
			fmt.Print(string(serverLine))
		case io.EOF:
			log.Fatal("No more input to read from connection")
		default:
			log.Fatal("Somthing wrong happend ", err)
		}
	}
}
