package room

import (
	"chat/framework/communicator"
	"chat/framework/message"
	"chat/framework/structs"
	"chat/impl/server/user"
	"strconv"
)

// Rooms represents rooms in the chat server
type Rooms map[structs.RoomID]Room

// Room a
type Room interface {
	RunRoom()
	AddComm(communicator.Communicator) error
}

// room represents a single room on the chat server.
type room struct {
	users  []user.User
	roomCh chan message.Message
}

// NewRoom inits new room
func NewRoom() Room {
	return &room{
		users:  []user.User{},
		roomCh: make(chan message.Message),
	}
}

// RunRoom open the given room to recieve messages
func (room *room) RunRoom() {
	for {
		msg := <-room.roomCh
		room.publishMsg(msg)
	}
}

// publishMsg publish message to all the users in the room
func (room *room) publishMsg(msg message.Message) {
	for _, user := range room.users {
		if msg.User() != user.ID() {
			// It's much safer to copy the msg and not just passing it around
			// By now I couldn't find a good way to do it
			user.Send(msg)
		}
	}
}

// AddComm adds a given communicator to the room
func (room *room) AddComm(comm communicator.Communicator) error {
	id := structs.User{
		ID:   strconv.Itoa(len(room.users)),
		Name: ""}

	user, err := user.NewUser(id, comm, room.roomCh)
	if err != nil {
		return err
	}

	room.users = append(room.users, user)
	go user.Serve()
	return nil
}
