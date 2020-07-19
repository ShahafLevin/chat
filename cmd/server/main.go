package main

import (
	"chat/impl/server"
	"flag"
)

var (
	port = flag.Int("port", 0, "The server port")
)

func main() {
	flag.Parse()
	tcpServer := server.NewServer(*port, "tcp")
	tcpServer.Run()
}
