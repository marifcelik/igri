import { useContext, useEffect, useRef } from 'react'
import { useAutoAnimate } from '@formkit/auto-animate/react'
import { ChatContext } from '@/context/chatContext'
import Bubble from './Bubble'
import { UserContext } from '@/context/userContext'

export default function ChatContainer() {
	const [autoAnimateRef] = useAutoAnimate<HTMLDivElement>({ duration: 100, easing: 'linear' })

	const chatContainer = useRef<HTMLDivElement | null>(null)
	const { messageHistory, receiver } = useContext(ChatContext)!
	const { user } = useContext(UserContext)!

	useEffect(() => {
		if (chatContainer.current)
			chatContainer.current.scrollTo({ top: chatContainer.current.scrollHeight, behavior: 'smooth' })
	}, [messageHistory, chatContainer])

	return receiver === '' || receiver === '""' ? (
		<div className="h-[calc(100%-8rem)] flex items-center justify-center">
			<div className="text-muted">Select a chat to start messaging</div>
		</div>
	) : (
		<div ref={chatContainer} className="h-[calc(100%-8rem)] overflow-y-scroll p-5">
			<div ref={autoAnimateRef} className="h-full">
				{messageHistory.map((message, i) => (
					<Bubble key={i} position={message.senderID === user.id ? 'right' : 'left'}>
						{message.data}
					</Bubble>
				))}
			</div>
		</div>
	)
}
