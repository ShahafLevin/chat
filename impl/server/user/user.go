package user

import (
	"chat/framework/communicator"
	"chat/framework/message"
	"chat/framework/structs"
)

// User in the chat server
type User interface {
	Serve()
	ID() structs.User
	Send(message.Message)
}

// user in the chat server
type user struct {
	id        structs.User
	comm      communicator.Communicator
	sendCh    chan message.Message
	recieveCh chan message.Message
}

// NewUser creates new user
func NewUser(id structs.User, comm communicator.Communicator, sendCh chan message.Message) (User, error) {
	return &user{
		id:        id,
		comm:      comm,
		sendCh:    sendCh,
		recieveCh: make(chan message.Message)}, nil
}

// Serve handels a user requests and response
func (user *user) Serve() {
	go user.comm.Recieve(user.sendCh)
	user.comm.Send(user.recieveCh)
}

func (user *user) ID() structs.User {
	return user.id
}

func (user *user) Send(msg message.Message) {
	user.recieveCh <- msg
}
