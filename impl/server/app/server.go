package app

import (
	"chat/framework/cryptochat"
	"chat/impl/server/connector"
	"chat/impl/server/room"
	"log"
	"net"
)

// Server represents a chat server
type Server interface {
	Run()
}

type server struct {
	rooms     room.Rooms
	connector connector.Connector
}

// NewServer init a new server
func NewServer(conn connector.Connector) (*Server, error) {
	server := Server{
		rooms:     InitRooms(),
		connector: conn,
	}
	return &server, nil
}

// Run is the method which runs the server
func (s *server) Run() error {
	// TODO: implement all the connection logic as  part of the connector
	// ln, err := net.Listen("tcp", fmt.Sprintf(":%d", server.port))

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Listening at port: %d", server.port)
	// server.listener = ln
	// defer server.listener.Close()

	if err := s.connector.Listen(); err != nil {
		return err
	}
	return nil
	//s.handleUsers(s.connector.Listen)
}

// handleUsers handle the users that connects to the chat room
// Note: probably will be removed after communicator refactor
func (s *server) handleUsers() {
	for {
		// TODO: implement all the connection logic as  part of the connector
		// conn, err := server.listener.Accept()
		// if err != nil {
		// 	log.Println(err)
		// 	continue
		// }
		// defer conn.Close()

		// secret := establishSecret(conn)
		// roomAsBytes, err := bufio.NewReader(conn).ReadBytes(byte('\n'))
		// if err != nil {
		// 	log.Println(err)
		// 	conn.Close()
		// 	continue
		// }

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
// TODO: Load config from JSON, and pass the rooms by the ctor
func InitRooms() room.Rooms {
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
