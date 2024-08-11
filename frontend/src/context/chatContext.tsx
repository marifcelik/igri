import { createContext, useState } from 'react'
import type { WSMessage } from '@/types'

type ChatContextType = {
	messageHistory: WSMessage[]
	setMessageHistory: React.Dispatch<React.SetStateAction<WSMessage[]>>
}

export const ChatContext = createContext<ChatContextType | undefined>(undefined)

export function ChatProvider({ children }: { children: React.ReactNode }) {
	const [messageHistory, setMessageHistory] = useState<WSMessage[]>([])

	return <ChatContext.Provider value={{ messageHistory, setMessageHistory }}>{children}</ChatContext.Provider>
}
