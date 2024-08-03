package ws

import "go-chat/enums"

type MessageDTO struct {
	Type  enums.MessageType `json:"type"`
	From  string            `json:"from"`
	To    string            `json:"to,omitempty"`
	Group string            `json:"group,omitempty"`
	Data  string            `json:"data"`
}
