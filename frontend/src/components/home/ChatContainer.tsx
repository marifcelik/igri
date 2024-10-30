import { useContext, useEffect, useRef } from 'react'
import { useAutoAnimate } from '@formkit/auto-animate/react'
import { toast } from 'sonner'
import { ChatContext } from '@/context/chatContext'
import { UserContext } from '@/context/userContext'
import Bubble from './Bubble'
import type { WSMessage } from '@/types'
import { API_URL } from '@/lib/config'

export default function ChatContainer() {
	const [autoAnimateRef] = useAutoAnimate<HTMLDivElement>({ duration: 100, easing: 'linear' })

	const chatContainer = useRef<HTMLDivElement | null>(null)
	const { messageHistory, setMessageHistory, conversation } = useContext(ChatContext)!
	const { user } = useContext(UserContext)!

	useEffect(() => {
		if (conversation?.id) {
			setMessageHistory([])
			fetchMessages()
		}
	}, [conversation])

	useEffect(() => {
		if (chatContainer.current)
			chatContainer.current.scrollTo({ top: chatContainer.current.scrollHeight, behavior: 'smooth' })
	}, [messageHistory, chatContainer])

	async function fetchMessages() {
		try {
			const res = await fetch(`${API_URL}/messages/${conversation!.id}`, {
				headers: { 'X-Session': user.token }
			})
			if (!res.ok) throw new Error(res.statusText)

			const { data } = (await res.json()) as { data: WSMessage[] }
			setMessageHistory(data)
		} catch (err: any) {
			console.error(err)
			toast.error('Failed to fetch messages', { description: err?.message })
		}
	}

	return (
		<div ref={chatContainer} className="h-[calc(100%-8rem)] overflow-y-scroll">
			<div ref={autoAnimateRef} className="h-fit p-4 sm:p-2">
				{messageHistory.map((message, i) => (
					<Bubble key={i} position={message.senderID === user.id ? 'right' : 'left'}>
						{message.content}
					</Bubble>
				))}
			</div>
		</div>
	)
}
