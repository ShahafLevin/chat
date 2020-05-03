package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	address := flag.String("address", "", "The address for the chat room as ip:port")
	room := flag.Int("room", 0, "The room number")

	flag.Parse()

	if *address == "" {
		log.Fatal("Address must be provided!")
	}

	if *room == 0 {
		log.Fatal("Room number must be provided!")
	}

	log.Printf("Connecting to room %d at %s \n", *room, *address)
	conn, err := net.Dial("tcp", *address)

	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}

	defer conn.Close()

	connectToRoom(conn, *room)

	go userInput(conn)
	serverResponse(conn)
}

// connectToRoom connects to the server with the given room number
func connectToRoom(conn net.Conn, room int) {
	conn.Write([]byte(strconv.Itoa(room) + "\n"))
	repsonse, err := bufio.NewReader(conn).ReadByte()

	if err != nil {
		log.Fatal(err)
	}

	if repsonse == '1' {
		log.Fatal("Room does not exist")
	}

	log.Println("Connected!")

}

// userInput reads the user input and sends it to the connection
func userInput(conn net.Conn) {
	userInput := bufio.NewReader(os.Stdin)
	for {
		userLine, err := userInput.ReadBytes(byte('\n'))
		switch err {
		case nil:
			conn.Write(userLine)
		case io.EOF:
			log.Fatal("No more output to send to connection")
		default:
			log.Fatal("Somthing wrong happend ", err)
		}
	}
}

// serverResponse recieves from the server a message and prints it
func serverResponse(conn net.Conn) {
	response := bufio.NewReader(conn)
	for {
		serverLine, err := response.ReadBytes(byte('\n'))
		switch err {
		case nil:
			fmt.Print(string(serverLine))
		case io.EOF:
			log.Fatal("No more input to read from connection")
		default:
			log.Fatal("Somthing wrong happend ", err)
		}
	}
}
