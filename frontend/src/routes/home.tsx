import { useState, useEffect, useRef, type FormEvent } from 'react'
import useWebSocket from 'react-use-websocket'
import { createFileRoute, redirect, Link } from '@tanstack/react-router'
import { PhoneIcon, SendIcon, VideoIcon } from 'lucide-react'
import { useAutoAnimate } from '@formkit/auto-animate/react'
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import Sidebar from '@/components/home/Sidebar'
import Bubble from '@/components/home/Bubble'
import ThemeButton from '@/components/ThemeButton'

export const Route = createFileRoute('/home')({
	beforeLoad: async ({ location }) => {
		const token = localStorage.getItem('token')
		console.log(token)
		if (!token || token === 'null' || token === 'undefined')
			throw redirect({
				to: '/login',
				search: { redirect: location.href }
			})
	},
	component: Home
})

type WSMessage = {
	type: 0 | 1
	from: string
	to: string
	data: string
	group?: string
}

// TODO redesign this page, seperate the chat list and chat content into two different components
function Home() {
	const [messages, setMessages] = useState<{ message: string; position: 'left' | 'right' }[]>([
		{ message: 'Hey, how are you?', position: 'right' },
		{ message: "I'm good, thanks! How about you?", position: 'left' },
		{ message: "I'm doing great, thanks! Just got back from a trip.", position: 'right' },
		{ message: 'That sounds amazing! Where did you go?', position: 'left' },
		{ message: 'I went to the beach with some friends.', position: 'right' },
		{ message: "That sounds like so much fun! I'm jealous.", position: 'left' },
		{ message: 'You should come next time! We had a blast.', position: 'right' },
		{ message: "I'd love to! Let me know when you're planning the next trip.", position: 'left' },
		{ message: "Will do! So, what's new with you?", position: 'right' },
		{ message: 'Not much, just work and stuff. You know how it is.', position: 'left' },
		{ message: 'Yeah, I do. It can be tough sometimes.', position: 'right' },
		{ message: 'Definitely. But hey, at least we have the weekends, right?', position: 'left' },
		{ message: 'Exactly! The weekends are the best.', position: 'right' },
		{ message: 'So, do you have any fun plans for this weekend?', position: 'left' },
		{ message: 'Actually, I was thinking of checking out that new restaurant downtown.', position: 'right' },
		{ message: "I've been meaning to try that place out too!", position: 'left' },
		{ message: "Yeah, I've heard great things about it.", position: 'left' },
		{ message: 'Definitely let me know how it is if you end up going!', position: 'right' },
		{ message: "Will do! I'm actually thinking of going tomorrow night.", position: 'right' },
		{ message: 'Cool, have fun!', position: 'left' },
		{ message: 'Thanks, I will!', position: 'right' },
		{ message: "So, how's your family doing?", position: 'left' },
		{ message: "They're all good, thanks for asking!", position: 'right' },
		{ message: "My mom's been bugging me to come visit again.", position: 'right' },
		{ message: "I'm sure she misses you! You should go see her soon.", position: 'left' }
	])
	const [messageValue, setMessageValue] = useState('')
	const [disabled, setDisabled] = useState(true)
	const [messageHistory, setMessageHistory] = useState<WSMessage[]>([])
	const chatContainer = useRef<HTMLDivElement | null>(null)

	const [autoAnimateRef] = useAutoAnimate<HTMLDivElement>()

	const username = localStorage.getItem('username') as string
	const to = localStorage.getItem('to') as string
	const token = localStorage.getItem('token') as string

	useEffect(() => {
		if (chatContainer.current)
			chatContainer.current.scrollTo({ top: chatContainer.current.scrollHeight, behavior: 'smooth' })
	}, [messageHistory, messages, chatContainer])

	useEffect(() => {
		setDisabled(messageValue.length === 0)
	}, [messageValue])

	const { sendJsonMessage, lastJsonMessage } = useWebSocket<WSMessage>('ws://localhost:8080/_ws?token=' + token, {
		onOpen: () => console.log('opened'),
		onClose: () => console.log('closed'),
		shouldReconnect: () => true
	})

	useEffect(() => {
		if (lastJsonMessage) {
			setMessageHistory(prev => prev.concat(lastJsonMessage))
		}
	}, [lastJsonMessage])

	async function handleSend(e: FormEvent<HTMLFormElement>) {
		try {
			e.preventDefault()
			console.log(messageValue)
			const msg: WSMessage = { type: 0, from: username, to: to, data: messageValue }
			sendJsonMessage<WSMessage>(msg)
			setMessageHistory([...messageHistory, msg])
			setMessageValue('')
		} catch (err) {
			// TODO handle error
			console.error(err)
		}
	}

	return (
		<div className="w-full h-full">
			<div className="grid grid-cols-[2fr_5fr] w-full h-96 lg:h-4/5 lg:w-9/12 absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 rounded-2xl border">
				<Sidebar />
				<div className="h-full overflow-hidden">
					<div className="p-3 flex border-b items-center h-16">
						<div className="flex items-center gap-2">
							<Avatar className="border w-10 h-10">
								<AvatarImage src="/placeholder-user.jpg" />
								<AvatarFallback>OM</AvatarFallback>
							</Avatar>
							<div className="grid gap-0.5">
								<p className="text-sm font-medium leading-none">Sofia Davis</p>
								<p className="text-xs text-muted-foreground">Active 2h ago</p>
							</div>
						</div>
						<div className="flex items-center gap-1 ml-auto">
							<Button variant="ghost" size="icon">
								<span className="sr-only">Call</span>
								<PhoneIcon className="h-4 w-4" />
							</Button>
							<Button variant="ghost" size="icon">
								<span className="sr-only">Video call</span>
								<VideoIcon className="h-4 w-4" />
							</Button>
							<Button variant="ghost" asChild>
								<Link to="/login">Login</Link>
							</Button>
							<ThemeButton />
						</div>
					</div>
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
					<div id="messageBox" className="border-t h-16 w-full">
						<form onSubmit={handleSend} className="flex w-full items-center space-x-2 p-3">
							<Input
								name="message"
								placeholder="Type your message..."
								className="flex-1"
								autoComplete="off"
								value={messageValue}
								onChange={e => setMessageValue(e.target.value)}
							/>
							<Button type="submit" size="icon" disabled={disabled}>
								<span className="sr-only">Send</span>
								<SendIcon className="h-4 w-4" />
							</Button>
						</form>
					</div>
				</div>
			</div>
		</div>
	)
}
