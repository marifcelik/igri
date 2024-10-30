package message

type MessageDTO struct {
	ID             string `json:"id"`
	Content        string `json:"content"`
	SenderID       string `json:"senderID"`
	CreatedAt      string `json:"createdAt"`
	ConversationID string `json:"conversationID,omitempty"`
}

type ConversationDTO struct {
	ID           string     `json:"id"`
	Name         string     `json:"name,omitempty"`
	Participants []string   `json:"participants"`
	LastMessage  MessageDTO `json:"lastMessage"`
}
