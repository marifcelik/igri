import { createContext } from 'react'
import type { ConversationMessage, ConversationPreview, ConversationRequest } from '@/types'

type ChatContextType = {
	messageHistory: ConversationMessage[]
	setMessageHistory: React.Dispatch<React.SetStateAction<ConversationMessage[]>>
	recipientID: string | null
	setRecipientID: React.Dispatch<React.SetStateAction<string | null>>
	conversation: ConversationPreview | null
	setConversation: React.Dispatch<React.SetStateAction<ConversationPreview | ConversationRequest | null>>
}

export const ChatContext = createContext<ChatContextType | undefined>(undefined)
