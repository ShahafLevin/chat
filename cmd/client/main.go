package main

import (
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
	tcpClient := NewClient(*address, *room, "tcp")
	tcpClient.Run()

}
