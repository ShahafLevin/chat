package app

import (
	"chat/framework/structs"
	"chat/impl/server/connector"
	"chat/impl/server/room"
)

// Server represents a chat server
type Server interface {
	Run() error
}

type server struct {
	rooms     room.Rooms
	connector connector.Connector
}

// NewServer init a new server
func NewServer(conn connector.Connector) Server {
	return &server{
		rooms:     InitRooms(),
		connector: conn,
	}
}

// Run is the method which runs the server
func (s *server) Run() error {
	if err := s.connector.Serve(s.rooms); err != nil {
		return err
	}
	return nil
}

// InitRooms initiate the rooms on the chat server
// TODO: Load config from JSON, and pass the rooms by the ctor
func InitRooms() room.Rooms {
	rooms := map[structs.RoomID]room.Room{
		structs.RoomID("1"): room.NewRoom(),
		structs.RoomID("2"): room.NewRoom(),
		structs.RoomID("3"): room.NewRoom(),
	}

	for _, room := range rooms {
		go room.RunRoom()
	}

	return rooms
}
