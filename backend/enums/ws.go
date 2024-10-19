package enums

type MessageResultStatus int

const (
	MessageResultSuccess MessageResultStatus = iota
	MessageResultFailed
)

func (e MessageResultStatus) String() string {
	switch e {
	case MessageResultSuccess:
		return "Success"
	case MessageResultFailed:
		return "Failed"
	default:
		return "Unknown"
	}
}

type WSMessageType int

const (
	WSMessageConversation WSMessageType = iota
	WSMessageResult
)

func (e WSMessageType) String() string {
	switch e {
	case WSMessageConversation:
		return "Conversation"
	case WSMessageResult:
		return "Result"
	default:
		return "Unknown"
	}
}
