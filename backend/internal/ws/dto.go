package ws

import "go-chat/enums"

type WsMessageDTO struct {
	Type     enums.MessageType `json:"type"`
	Message  string            `json:"message"`
	Sender   string            `json:"sender"`
	Receiver string            `json:"receiver,omitempty"`
	Group    string            `json:"group,omitempty"`
}
