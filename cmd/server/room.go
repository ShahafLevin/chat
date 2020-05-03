package main

import (
	"net"
	"strconv"
)

// RoomID represents the room ID
type RoomID string

// Room represents a single room on the chat server.
type Room struct {
	clients []Client
	roomCh  chan Message
}

// RunRoom open the given room to recieve messages
func (room *Room) RunRoom() {
	for {
		msg := <-room.roomCh
		room.publishMsg(msg)
	}
}

// publishMsg publish message to all the clients in the room
func (room *Room) publishMsg(msg Message) {
	for _, client := range room.clients {
		if msg.sender != client.id {
			client.recieveCh <- msg
		}
	}
}

// AddConn adds a given connection to the room
func (room *Room) AddConn(conn net.Conn) {
	client := Client{
		id:        ClientID(strconv.Itoa(len(room.clients))),
		conn:      conn,
		sendCh:    room.roomCh,
		recieveCh: make(chan Message)}

	room.clients = append(room.clients, client)
	client.Serve()
}
