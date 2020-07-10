package main

import (
	"flag"
)

var (
	port = flag.Int("port", 0, "The server port")
)

func main() {
	flag.Parse()
	tcpServer := NewServer(*port, "tcp")
	tcpServer.Run()
}
