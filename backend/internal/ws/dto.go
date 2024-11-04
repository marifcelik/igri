package ws

import "go-chat/enums"

type WSMessage struct {
	Type enums.WSMessageType `json:"type"`
	Data any                 `json:"data"`
}

type MessageDTO struct {
	Type              enums.ConversationType `json:"type"`
	SenderID          string                 `json:"senderID"`
	RecipientUsername string                 `json:"recipientUsername,omitempty"`
	ConversationID    string                 `json:"conversationID,omitempty"`
	Content           string                 `json:"content"`
	CreatedAt         string                 `json:"createdAt"`
}

type ResultDTO struct {
	Status         enums.ResultStatus `json:"status"`
	ConversationID string             `json:"conversationID"`
	MessageID      string             `json:"messageID"`
	Message        string             `json:"message,omitempty"`
}
