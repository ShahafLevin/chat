package main

import (
	"net"
	"strconv"
)

// Room represents a single room on the chat server.
type Room struct {
	users  []User
	roomCh chan Message
}

// RunRoom open the given room to recieve messages
func (room *Room) RunRoom() {
	for {
		msg := <-room.roomCh
		room.publishMsg(msg)
	}
}

// publishMsg publish message to all the users in the room
func (room *Room) publishMsg(msg Message) {
	for _, user := range room.users {
		if msg.sender != user.id {
			user.recieveCh <- msg
		}
	}
}

// AddConn adds a given connection to the room
func (room *Room) AddConn(conn net.Conn) {
	user := User{
		id:        UserID(strconv.Itoa(len(room.users))),
		conn:      conn,
		sendCh:    room.roomCh,
		recieveCh: make(chan Message)}

	room.users = append(room.users, user)
	user.Serve()
}
