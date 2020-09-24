package main

import (
	"chat/app"
	"chat/impl/server/connector"
	"flag"
	"os"
)

var (
	port   = flag.Int("port", 8080, "The server port")
	method = flag.String("method", "tcp", "The method in which the server use in order serve clients (tcp or udp)")
)

func main() {
	flag.Parse()

	tcpConnector := connector.NewNetConnector(*port, *method)

	tcpServer := app.NewServer(tcpConnector)
	if err := tcpServer.Run(); err != nil {
		os.Exit(1)
	}
}
