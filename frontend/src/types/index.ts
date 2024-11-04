type RequireField<T, K extends keyof T> = T & Required<Pick<T, K>>

/* TODO find another name for MessageType and generate another enum 
for the "actual" message type like text, image, video, etc. */
export enum MessageType {
	NORMAL = 0,
	GROUP
}

export type ConversationMessage = {
	type: MessageType
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

export enum ResultStatus {
	SUCCESS = 0,
	FAILURE
}

export type ResultMessage = {
	status: ResultStatus
	message?: string
	conversationID?: string
	messageID?: string
}

export enum WSMessageType {
	CONVERSATION = 0,
	RESULT
}

export type WSMessage =
	| { type: WSMessageType.CONVERSATION; data: ConversationMessage }
	| { type: WSMessageType.RESULT; data: ResultMessage }

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
	lastMessage: RequireField<Omit<ConversationMessage, 'recipientUsername'>, 'createdAt'>
	unreadCount?: number
}

export type ConversationRequest = {
	username: string
}
