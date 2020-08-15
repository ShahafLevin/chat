package main

import (
	"chat/framework/structs"
	"chat/impl/client/app"
	"chat/impl/client/connector"
	"flag"
	"os"
)

var (
	address = flag.String("address", "localhost:8080", "The address for the chat room as ip:port")
	room    = flag.String("room", "1", "The room number")
	method  = flag.String("method", "tcp", "The method to use in order to connect the room (tcp or udp)")
)

func main() {
	flag.Parse()

	tcpConnector, err := connector.NewNetConnector(*address, *method)
	if err != nil {
		os.Exit(1)
	}

	tcpClient, err := app.NewClient(tcpConnector, structs.RoomID(*room))
	if err != nil {
		os.Exit(1)
	}

	tcpClient.Run()
}
