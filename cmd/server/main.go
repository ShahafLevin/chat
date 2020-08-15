package main

import (
	"chat/impl/client/connector"
	"chat/impl/server/app"
	"flag"
	"os"
)

var (
	port   = flag.Int("port", 8080, "The server port")
	method = flag.String("method", "tcp", "The method in which the server use in order serve clients (tcp or udp)")
)

func main() {
	flag.Parse()

	tcpConnector, err := connector.NewNetConnector(*port, *method)
	if err != nil {
		os.Exit(1)
	}

	tcpServer, err := app.NewServer(connector)
	if err != nil {
		os.Exit(1)
	}

	tcpServer.Run()
}
