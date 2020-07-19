package server

import (
	"bufio"
	"chat/framework/cryptochat"
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
	key       cryptochat.Key
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
		msg, err := connReader.ReadBytes(byte('\n'))

		if err != nil {
			break
		}

		user.sendCh <- Message{
			content: cryptochat.DecryptMessage(user.key, msg[:len(msg)-1]),
			sender:  user.id}
	}
}

// recieveFromRoom recieves a message from the room and send it to the user
func (user *User) recieveFromRoom() {
	for {
		msg := <-user.recieveCh
		user.conn.Write(cryptochat.EncryptMessage(user.key, msg.content))
		user.conn.Write([]byte("\n"))
	}
}
