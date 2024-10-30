package enums

//go:generate enumer -type=ConversationType -trimsuffix=Conversation
type ConversationType int

// TODO support these message types
// const (
// 	MessageType_Text MessageType = iota
// 	MessageType_Image
// 	MessageType_Video
// 	MessageType_Audio
// 	MessageType_File
// )

const (
	NormalConversation ConversationType = iota
	GroupConversation
)
