package ws

import "go-chat/enums"

type WSMessage struct {
	Type enums.WSMessageType `json:"type"`
	Data any                 `json:"data"`
}

type ConversationMessageDTO struct {
	Type           enums.ConversationType `json:"type"`
	SenderID       string                 `json:"senderID"`
	ConversationID string                 `json:"conversationID,omitempty"`
	Content        string                 `json:"content"`
}

type MessageResultDTO struct {
	Status         enums.MessageResultStatus `json:"status"`
	ConversationID string                    `json:"conversationID"`
	MessageID      string                    `json:"messageID"`
}
