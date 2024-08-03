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
import { API_URL } from '@/lib/config'

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
	}, [messageHistory, chatContainer])

	useEffect(() => {
		setDisabled(messageValue.length === 0)
	}, [messageValue])

	const { sendJsonMessage, lastJsonMessage } = useWebSocket<WSMessage>('wss://' + API_URL + '/_ws?token=' + token, {
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
