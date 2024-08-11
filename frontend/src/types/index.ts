export enum MessageType {
	NORMAL = 0,
	GROUP
}

/* TODO find another name for MessageType and generate another enum 
for the "actual" message type like text, image, video, etc. */

export type WSMessage = {
	type: MessageType
	from: string
	to?: string
	data: string
	group?: string
}

export type UserFields = {
  id: string
  username: string
  token: string
}