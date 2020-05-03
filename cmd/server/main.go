package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 0, "The server port")

	flag.Parse()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Listening at port: %d", *port)
	}
	defer ln.Close()

	rooms := initRooms()

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Println(err)
			continue
		}
		defer conn.Close()

		roomAsBytes, err := bufio.NewReader(conn).ReadBytes(byte('\n'))
		if err != nil {
			log.Println(err)
			conn.Close()
			continue
		}

		roomID := RoomID(roomAsBytes[:len(roomAsBytes)-1])
		room, ok := rooms[roomID]

		if ok == false {
			conn.Write([]byte{'1'})
			conn.Close()
			continue
		} else {
			conn.Write([]byte{'2'})
		}

		room.AddConn(conn)
		log.Println("New client connected to Room", roomID)
	}
}

func initRooms() map[RoomID]*Room {
	rooms := map[RoomID]*Room{
		RoomID("1"): &Room{clients: []Client{}, roomCh: make(chan Message)},
		RoomID("2"): &Room{clients: []Client{}, roomCh: make(chan Message)},
		RoomID("3"): &Room{clients: []Client{}, roomCh: make(chan Message)},
	}

	for _, room := range rooms {
		go room.RunRoom()
	}

	return rooms
}
