package main

import (
	"flag"
)

func main() {
	port := flag.Int("port", 0, "The server port")
	flag.Parse()

	tcpServer := Server{
		port:   *port,
		rooms:  InitRooms(),
		method: "tcp",
	}

	tcpServer.Run()
}
