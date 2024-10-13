package ws

import "go-chat/enums"

type MessageDTO struct {
	Type       enums.MessageType `json:"type"`
	SenderID   string            `json:"senderID"`
	ReceiverID string            `json:"receiverID,omitempty"`
	GroupID    string            `json:"groupID,omitempty"`
	Data       string            `json:"data"`
}
