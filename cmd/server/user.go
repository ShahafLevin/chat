package main

import (
	"bufio"
	"net"
)

// UserID is the id of the user
type UserID string

// User in the chat server
type User struct {
	id        UserID
	conn      net.Conn
	sendCh    chan Message
	recieveCh chan Message
}

// Serve handels a user requests and response
func (user *User) Serve() {
	go user.sendToRoom()
	go user.recieveFromRoom()
}

// sendToRoom sends a message from the user to the room
func (user *User) sendToRoom() {
	connReader := bufio.NewReader(user.conn)
	for {
		msg, _ := connReader.ReadBytes(byte('\n'))

		user.sendCh <- Message{
			content: msg,
			sender:  user.id}
	}
}

// recieveFromRoom recieves a message from the room and send it to the user
func (user *User) recieveFromRoom() {
	for {
		msg := <-user.recieveCh
		user.conn.Write(msg.content)
	}
}
