package main

import (
	"chat/impl/client"
	"flag"
	"log"
)

var (
	address = flag.String("address", "localhost:8080", "The address for the chat room as ip:port")
	room    = flag.String("room", "1", "The room number")
)

func main() {
	flag.Parse()
	if *address == "" {
		log.Fatal("Address must be provided!")
	}
	if *room == "" {
		log.Fatal("Room number must be provided!")
	}

	tcpConnector, _ := client.NewNetConnector(*address, "tcp", false)
	tcpClient := client.Client{
		Connector: tcpConnector,
		Room:      *room,
	}
	tcpClient.Run()

}
