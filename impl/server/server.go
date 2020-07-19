package server

import (
	"bufio"
	"chat/framework/cryptochat"
	"fmt"
	"log"
	"net"
)

// RoomID represents the room ID
type RoomID string

type chatServer interface {
	Run()
	handleUsers()
}

// Server - a struct for the chat server
type Server struct {
	port     int
	rooms    map[RoomID]*Room
	method   string
	listener net.Listener
}

// NewServer init a new server
func NewServer(port int, method string) *Server {
	server := Server{
		port:   port,
		rooms:  InitRooms(),
		method: "tcp",
	}
	return &server
}

// Run is the method which runs the server
func (server *Server) Run() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", server.port))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening at port: %d", server.port)
	server.listener = ln
	defer server.listener.Close()

	server.handleUsers()
}

// handleUsers handle the users that connects to the chat room
func (server *Server) handleUsers() {
	for {
		conn, err := server.listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}
		defer conn.Close()

		secret := establishSecret(conn)
		roomAsBytes, err := bufio.NewReader(conn).ReadBytes(byte('\n'))
		if err != nil {
			log.Println(err)
			conn.Close()
			continue
		}

		roomID := RoomID(roomAsBytes[:len(roomAsBytes)-1])
		room, ok := server.rooms[roomID]

		if ok == false {
			conn.Write([]byte{'1'})
			conn.Close()
			continue
		}

		conn.Write([]byte{'2'})
		room.AddConn(conn, secret)
		log.Println("New user connected to Room", roomID)
	}
}

// InitRooms initiate the rooms on the chat server
func InitRooms() map[RoomID]*Room {
	rooms := map[RoomID]*Room{
		RoomID("1"): &Room{users: []User{}, roomCh: make(chan Message)},
		RoomID("2"): &Room{users: []User{}, roomCh: make(chan Message)},
		RoomID("3"): &Room{users: []User{}, roomCh: make(chan Message)},
	}

	for _, room := range rooms {
		go room.RunRoom()
	}

	return rooms
}

// establishSecret establish secret with the user
func establishSecret(conn net.Conn) (secert []byte) {
	key := cryptochat.GenerateKey()
	cryptochat.WriteKey(conn, (*key))
	userKey := cryptochat.ReadKey(conn, (*key))
	return key.KeyExchange.ComputeSecret(key.Private, userKey)
}
