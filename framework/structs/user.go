package structs

import (
	"fmt"

	"github.com/google/uuid"
)

// User represents a user in the chat
type User struct {
	ID   string
	Name string
}

// New User creates a new user using a given name
func NewUser(name string) (*User, error) {
	id, err := createUserId(name)
	if err != nil {
		return nil, err
	}
	return &User{ID: id, Name: name}, nil
}

// createUserId creates user id in the chat server
func createUserId(name string) (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s_%s", name, uuid), nil
}
