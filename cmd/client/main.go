package main

import (
	"chat/app"
	"chat/framework/structs"
	"chat/impl/client/connector"
	"flag"
	"os"
)

var (
	address = flag.String("address", "localhost:8080", "The address for the chat room as ip:port")
	room    = flag.String("room", "1", "The room number")
	method  = flag.String("method", "tcp", "The method to use in order to connect the room (tcp or udp)")
	name    = flag.String("name", "anonymous", "Enter your username")
)

func main() {
	flag.Parse()

	user, err := structs.NewUser(*name)
	if err != nil {
		os.Exit(1)
	}
	tcpConnector := connector.NewNetConnector(*address, *method)

	tcpClient, err := app.NewClient(tcpConnector, structs.RoomID(*room), *user)
	if err != nil {
		os.Exit(1)
	}

	tcpClient.Run()
}
