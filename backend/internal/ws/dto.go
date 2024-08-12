package ws

import "go-chat/enums"

type MessageDTO struct {
	Type     enums.MessageType `json:"type"`
	Sender   string            `json:"sender"`
	Receiver string            `json:"receiver,omitempty"`
	Group    string            `json:"group,omitempty"`
	Data     string            `json:"data"`
}
