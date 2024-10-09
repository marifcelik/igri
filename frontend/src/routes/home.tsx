import { useState, useEffect, type FormEvent, useContext } from 'react'
import useWebSocket, { ReadyState } from 'react-use-websocket'
import { createFileRoute, redirect } from '@tanstack/react-router'
import { toast } from 'sonner'
import Sidebar from '@/components/home/Sidebar'
import PersonInfo from '@/components/home/PersonInfo'
import ChatContainer from '@/components/home/ChatContainer'
import MessageBox from '@/components/home/MessageBox'
import { API_URL } from '@/lib/config'
import { type WSMessage, MessageType } from '@/types'
import { UserContext } from '@/context/userContext'
import { ChatContext } from '@/context/chatContext'

export const Route = createFileRoute('/home')({
	beforeLoad: () => {
		const token = localStorage.getItem('token')
		console.log(token)
		if (!token || token === 'null' || token === 'undefined' || token === '""')
			throw redirect({
				to: '/login',
				search: { redirect: '/home' }
			})
	},
	component: Home
})

// TODO redesign this page, seperate the chat list and chat content into two different components
function Home() {
	const [messageValue, setMessageValue] = useState('')
	const { user } = useContext(UserContext)!
	const { setMessageHistory } = useContext(ChatContext)!

	const navigate = Route.useNavigate()

	// TODO find better place for to, maybe in chat context
	const { receiver } = useContext(ChatContext)!

	const { sendJsonMessage, lastJsonMessage, readyState } = useWebSocket<WSMessage>(
		API_URL.replace('http', 'ws') + '/_ws?token=' + user.token,
		{
			onOpen: () => console.log('opened'),
			onClose: () => console.log('closed'),
			shouldReconnect: () => true,
			onReconnectStop: () => {
				// TODO use `useLocalStorage` hook
				localStorage.setItem('token', '""')
				toast.error('Cannot connect to chat server, please login again', { duration: 5000 })
				navigate({ to: '/login', search: { redirect: '/home' } })
			},
			reconnectAttempts: 3,
			reconnectInterval: 2000
		}
	)

	useEffect(() => {
		toast.info('receiver is ' + receiver)
	}, [receiver])

	useEffect(() => {
		if (lastJsonMessage) {
			setMessageHistory(prev => prev.concat(lastJsonMessage))
			console.log('lastJsonMessage', lastJsonMessage)
		}
	}, [lastJsonMessage])

	async function handleSend(e: FormEvent<HTMLFormElement>) {
		if (!receiver) return

		// FIX sometimes it sends an empty message
		try {
			e.preventDefault()
			console.log('message value home', messageValue)
			console.log('receiver', receiver)
			const msg: WSMessage = {
				type: MessageType.NORMAL,
				senderID: user.username,
				receiverID: receiver,
				data: messageValue
			}
			sendJsonMessage<WSMessage>(msg)
			setMessageHistory(prev => prev.concat(msg))
			setMessageValue('')
		} catch (err) {
			// TODO handle error
			console.error(err)
		}
	}

	return (
		<div className="w-full h-full">
			<div className="grid grid-cols-[2fr_5fr] w-full h-96 lg:h-4/5 lg:w-9/12 absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 rounded-2xl border">
				{readyState === ReadyState.OPEN ? (
					<>
						<Sidebar />
						<div className="h-full overflow-hidden">
							<PersonInfo />
							<ChatContainer />
							<MessageBox onSubmit={handleSend} value={messageValue} setValue={setMessageValue} />
						</div>
					</>
				) : (
					// TODO i think there should be a better way to center text
					<div className="w-full h-full flex items-center justify-center">
						<div className="text-center">
							<h1 className="text-4xl font-bold">
								{readyState === ReadyState.CONNECTING ? 'Connecting...' : 'Disconnected'}
							</h1>
						</div>
					</div>
				)}
			</div>
		</div>
	)
}
