import { createContext } from 'react'
import type { WSMessage } from '@/types'

type ChatContextType = {
	messageHistory: WSMessage[]
	setMessageHistory: React.Dispatch<React.SetStateAction<WSMessage[]>>
	receiver: string | null
	setReceiver: React.Dispatch<React.SetStateAction<string | null>>
}

export const ChatContext = createContext<ChatContextType | undefined>(undefined)
