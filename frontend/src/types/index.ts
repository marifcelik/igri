type RequireField<T, K extends keyof T> = T & Required<Pick<T, K>>

/* TODO find another name for MessageType and generate another enum 
for the "actual" message type like text, image, video, etc. */
export enum MessageType {
	NORMAL = 0,
	GROUP
}

export type WSMessage = {
	id?: string
	senderID: string
	content: string
	createdAt?: string
} & (
	| {
			recipientUsername: string
			conversationID?: never
	  }
	| {
			recipientUsername?: never
			conversationID: string
	  }
)

export type UserFields = {
	id: string
	username: string
	token: string
}

export type ConversationPreview = {
	id: string
	name: string
	username: string
	participants: string[]
	image: string
	lastMessage: RequireField<Omit<WSMessage, 'recipientID'>, 'createdAt'>
}

export type ConversationRequest = {
	username: string
}
