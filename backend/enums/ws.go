package enums

//go:generate enumer -type=ResultStatus -trimprefix=Result
type ResultStatus int

const (
	ResultSuccess ResultStatus = iota
	ResultFailed
)

//go:generate enumer -type=WSMessageType -trimprefix=WSMessage
type WSMessageType int

const (
	WSMessageConversation WSMessageType = iota
	WSMessageResult
)
