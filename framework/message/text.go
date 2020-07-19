package message

import "chat/framework/user"

// TextType is the type of the text message
const TextType = "Text"

// Text represents a text message in the chat
type Text struct {
	BaseMessage
}

// NewText creates new text message
func NewText(content []byte, sender user.User) *Text {
	var text Text
	text = Text{BaseMessage{Content: content, Sender: sender, Type: TextType}}
	return &text
}

// Marshal marshals the text message
func (text *Text) Marshal() []byte {
	return text.Content
}

// UnMarshal unmarshals byte to our msg obj
func (text *Text) UnMarshal(buff []byte) {
	text.Content = buff
}
