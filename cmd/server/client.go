package main

import (
	"bufio"
	"net"
)

// ClientID is the id of the cliert
type ClientID string

// Client in the chat server
type Client struct {
	id        ClientID
	conn      net.Conn
	sendCh    chan Message
	recieveCh chan Message
}

// Serve handels a client requests and response
func (client *Client) Serve() {
	go client.sendToRoom()
	go client.recieveFromRoom()
}

// sendToRoom sends a message from the client to the room
func (client *Client) sendToRoom() {
	connReader := bufio.NewReader(client.conn)
	for {
		msg, _ := connReader.ReadBytes(byte('\n'))

		client.sendCh <- Message{
			content: msg,
			sender:  client.id}
	}
}

// recieveFromRoom recieves a message from the room and send it to the client
func (client *Client) recieveFromRoom() {
	for {
		msg := <-client.recieveCh
		client.conn.Write(msg.content)
	}
}
