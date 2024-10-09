/* TODO find another name for MessageType and generate another enum 
for the "actual" message type like text, image, video, etc. */
export enum MessageType {
	NORMAL = 0,
	GROUP
}

export type WSMessage = {
	type: MessageType
	senderID: string
	receiverID?: string
	data: string
	group?: string
}

export type UserFields = {
	id: string
	username: string
	token: string
}

export type ChatPerson = {
	id: string
	name: string
	username: string
	image: string
	time: string
	message: string
}
