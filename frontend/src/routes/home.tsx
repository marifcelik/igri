import { useState, useEffect, type FormEvent } from 'react'
import useWebSocket from 'react-use-websocket'
import { createFileRoute, redirect } from '@tanstack/react-router'
import { useLocalStorage } from '@uidotdev/usehooks'
import Sidebar from '@/components/home/Sidebar'
import PersonInfo from '@/components/home/PersonInfo'
import ChatContainer from '@/components/home/ChatContainer'
import MessageBox from '@/components/home/MessageBox'
import { ChatProvider } from '@/context/chatContext'
import { useChat, useUser } from '@/hooks/context'
import { API_URL } from '@/lib/config'
import { type WSMessage, MessageType } from '@/types'
import { UserProvider } from '@/context/userContext'

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
	component: () => (
		<ChatProvider>
			<UserProvider>
				<Home />
			</UserProvider>
		</ChatProvider>
	)
})

// TODO redesign this page, seperate the chat list and chat content into two different components
function Home() {
	const [messageValue, setMessageValue] = useState('')
	const { setMessageHistory } = useChat()
	const { user } = useUser()

	// TODO find better place for to, maybe in chat context
	const [to] = useLocalStorage<string>('to')

	const { sendJsonMessage, lastJsonMessage } = useWebSocket<WSMessage>(
		API_URL.replace('http', 'ws') + '/_ws?token=' + user.token,
		{
			onOpen: () => console.log('opened'),
			onClose: () => console.log('closed'),
			shouldReconnect: () => true
		}
	)

	useEffect(() => {
		if (lastJsonMessage) {
			setMessageHistory(prev => prev.concat(lastJsonMessage))
		}
	}, [lastJsonMessage])

	async function handleSend(e: FormEvent<HTMLFormElement>) {
		// FIX sometimes it sends an empty message
		try {
			e.preventDefault()
			console.log(messageValue)
			const msg: WSMessage = { type: MessageType.NORMAL, from: user.username, to: to, data: messageValue }
			sendJsonMessage<WSMessage>(msg)
			setMessageHistory(prev => prev.concat(lastJsonMessage))
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
					<PersonInfo />
					<ChatContainer />
					<MessageBox onSubmit={handleSend} value={messageValue} setValue={setMessageValue} />
				</div>
			</div>
		</div>
	)
}
