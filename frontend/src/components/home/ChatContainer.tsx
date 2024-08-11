import { useEffect, useRef } from 'react'
import { useAutoAnimate } from '@formkit/auto-animate/react'
import useChat from '@/hooks/useChat'
import { useLocalStorage } from '@uidotdev/usehooks'
import Bubble from './Bubble'

export default function ChatContainer() {
	const [username] = useLocalStorage<string>('username')

	const chatContainer = useRef<HTMLDivElement | null>(null)
	const [autoAnimateRef] = useAutoAnimate<HTMLDivElement>()

	const { messageHistory } = useChat()

	useEffect(() => {
		if (chatContainer.current)
			chatContainer.current.scrollTo({ top: chatContainer.current.scrollHeight, behavior: 'smooth' })
	}, [messageHistory, chatContainer])

	return (
		<div ref={chatContainer} className="h-[calc(100%-8rem)] overflow-y-scroll p-5">
			<div ref={autoAnimateRef} className="h-full">
				{messageHistory.map((message, i) => {
					return (
						<Bubble key={i} position={message.from === username ? 'right' : 'left'}>
							{message.data}
						</Bubble>
					)
				})}
			</div>
		</div>
	)
}
