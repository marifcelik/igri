import { useContext, useEffect, useRef } from 'react'
import { useAutoAnimate } from '@formkit/auto-animate/react'
import { useLocalStorage } from '@uidotdev/usehooks'
import { ChatContext } from '@/context/chatContext'
import Bubble from './Bubble'

export default function ChatContainer() {
	const [username] = useLocalStorage<string>('username')
	const [autoAnimateRef] = useAutoAnimate<HTMLDivElement>({ duration: 100 })

	const chatContainer = useRef<HTMLDivElement | null>(null)
	const { messageHistory } = useContext(ChatContext)!

	useEffect(() => {
		if (chatContainer.current)
			chatContainer.current.scrollTo({ top: chatContainer.current.scrollHeight, behavior: 'smooth' })
	}, [messageHistory, chatContainer])

	return (
		<div ref={chatContainer} className="h-[calc(100%-8rem)] overflow-y-scroll p-5">
			<div ref={autoAnimateRef} className="h-full">
				{messageHistory.map((message, i) => (
					<Bubble key={i} position={message.sender === username ? 'right' : 'left'}>
						{message.data}
					</Bubble>
				))}
			</div>
		</div>
	)
}
