package main

import (
	"flag"
	"log"
)

func main() {
	address := flag.String("address", "", "The address for the chat room as ip:port")
	room := flag.String("room", "", "The room number")

	flag.Parse()

	if *address == "" {
		log.Fatal("Address must be provided!")
	}

	if *room == "" {
		log.Fatal("Room number must be provided!")
	}

	tcpClient := Client{
		address: *address,
		room:    RoomID(*room),
		method:  "tcp",
	}

	tcpClient.Run()

}
