package server

// Message represents a message in the chat server
type Message struct {
	content []byte
	sender  UserID
}
