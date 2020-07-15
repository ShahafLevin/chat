package main

import (
	"bufio"
	"chat/cmd/cryptochat"
	"crypto"
	"crypto/elliptic"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/aead/ecdh"
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
	key     cryptochat.Key
}

// NewClient creates new client instance
func NewClient(addr string, room string, method string) *Client {
	client := Client{
		address: addr,
		room:    RoomID(room),
		method:  method,
	}
	log.Printf("Connecting to room %s at %s \n", client.room, client.address)
	conn, err := net.Dial(client.method, client.address)

	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}

	client.conn = conn
	// defer client.conn.Close()

	client.key = cryptochat.Key(establishSecret(&client.conn))
	fmt.Print(client.key)
	return &client
}

// Run is the method which runs the client
func (client *Client) Run() {
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
			client.conn.Write(cryptochat.EncryptMessage(client.key, userLine))
			client.conn.Write([]byte("\n"))
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
			fmt.Print(cryptochat.DecryptMessage(client.key, serverLine[:len(serverLine)-1]))
		case io.EOF:
			log.Fatal("No more input to read from connection")
		default:
			log.Fatal("Somthing wrong happend ", err)
		}
	}
}

// establishSecret establish secret with the user
func establishSecret(conn *net.Conn) (secert []byte) {
	key := cryptochat.GenerateKey()
	var point = key.Public.(ecdh.Point)
	serverBuf, _ := bufio.NewReader(*conn).ReadBytes(byte('\n'))
	(*conn).Write(elliptic.Marshal(key.Curve, point.X, point.Y))
	(*conn).Write([]byte("\n"))

	var serverKey crypto.PublicKey
	x, y := elliptic.Unmarshal(key.Curve, serverBuf[:len(serverBuf)-1])
	serverKey = ecdh.Point{X: x, Y: y}

	return key.KeyExchange.ComputeSecret(key.Private, serverKey)
}
