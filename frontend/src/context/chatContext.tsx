import { createContext } from 'react'
import type { WSMessage, ConversationPreview, ConversationRequest } from '@/types'

type ChatContextType = {
	messageHistory: WSMessage[]
	setMessageHistory: React.Dispatch<React.SetStateAction<WSMessage[]>>
	recipientID: string | null
	setRecipientID: React.Dispatch<React.SetStateAction<string | null>>
	conversation: ConversationPreview | null
	setConversation: React.Dispatch<React.SetStateAction<ConversationPreview | ConversationRequest | null>>
}

export const ChatContext = createContext<ChatContextType | undefined>(undefined)
