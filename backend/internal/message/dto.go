package message

type MessageDTO struct {
	ID          string `json:"id"`
	IsGroup     bool   `json:"isGroup"`
	IsFromUser  bool   `json:"isFromUser"`
	Name        string `json:"name"`
	LastMessage string `json:"lastMessage"`
}
