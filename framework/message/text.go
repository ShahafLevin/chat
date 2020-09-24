package message

import (
	"chat/framework/structs"
)

// TextType is the type of the text message
const TextType = "Text"

// Text represents a text message in the chat
type Text struct {
	BaseMessage
}

// NewText creates new text message
func NewText(content []byte, sender structs.User) Message {
	return &Text{BaseMessage{Content: content, Sender: sender, Type: TextType}}
}

// Marshal marshals the text message
func (text *Text) Marshal() []byte {
	return append([]byte(nil), text.BaseMessage.Content...)
}

// UnMarshal unmarshals byte to our msg obj
func (text *Text) UnMarshal(buff []byte) {
	text.BaseMessage.Content = buff
}

// User returns the user
func (text *Text) User() structs.User {
	return text.BaseMessage.Sender
}
