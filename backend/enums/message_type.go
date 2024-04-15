package enums

import log "github.com/charmbracelet/log"

type MessageType int

// TODO support these message types
// const (
// 	MessageType_Text MessageType = iota
// 	MessageType_Image
// 	MessageType_Video
// 	MessageType_Audio
// 	MessageType_File
// 	MessageType_Location
// 	MessageType_Contact
// 	MessageType_Sticker
// 	MessageType_Gif
// )

const (
	Normal_MessageType MessageType = iota
	Group_MessageType
)

func (m MessageType) String() string {
	var words = [...]string{"Normal", "Group"}

	if m < 0 || int(m) > len(words) {
		// if i use m directly, it will cause infinite loop because of String() method call itself recursively
		// so convert it to its underlying type
		// if i use log directly, it will cause the same problem too, idk why
		log.WithPrefix("WS").Warn("Invalid message type", "sended message type", int(m))
		return "Invalid"
	}
	return words[m]
}
