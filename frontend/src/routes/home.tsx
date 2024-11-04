import { useState, useEffect, type FormEvent, useContext, useRef, useMemo } from 'react'
import useWebSocket, { ReadyState } from 'react-use-websocket'
import { createFileRoute, redirect } from '@tanstack/react-router'
import { useAutoAnimate } from '@formkit/auto-animate/react'
import { toast } from 'sonner'
import Sidebar from '@/components/home/Sidebar'
import PersonInfo from '@/components/home/PersonInfo'
import ChatContainer from '@/components/home/ChatContainer'
import MessageBox from '@/components/home/MessageBox'
import { API_URL } from '@/lib/config'
import useBreakpoint from '@/hooks/breakpoint'
import { ConversationPreview, ConversationMessage, WSMessage, WSMessageType, ResultStatus, MessageType } from '@/types'
import { UserContext } from '@/context/userContext'
import { ChatContext } from '@/context/chatContext'

export const Route = createFileRoute('/home')({
	beforeLoad: () => {
		const token = localStorage.getItem('token')
		console.log(token)
		if (!token || ['null', '"null"', 'undefined', '"undefined"', '""'].includes(token))
			throw redirect({
				to: '/auth/login',
				search: { redirect: '/home' }
			})
	},
	component: Home
})

// TODO redesign this page, seperate the chat list and chat content into two different components
function Home() {
	const [messageValue, setMessageValue] = useState('')
	const [conversations, setConversations] = useState<ConversationPreview[]>([])

	const conversationPreviews = useMemo(() => {
		return conversations.sort((a, b) => {
			const aTime = new Date(a.lastMessage.createdAt).getTime()
			const bTime = new Date(b.lastMessage.createdAt).getTime()
			return bTime - aTime
		})
	}, [conversations])

	const inputRef = useRef<HTMLInputElement>(null)

	const { user, setUser } = useContext(UserContext)!
	const { setMessageHistory, conversation, setConversation } = useContext(ChatContext)!

	const [autoAnimateRef] = useAutoAnimate<HTMLDivElement>()

	const breakpoint = useBreakpoint()

	// TODO block back to the login page
	// useBlocker()
	const navigate = Route.useNavigate()

	const { sendJsonMessage, readyState } = useWebSocket<WSMessage>(
		API_URL.replace('http', 'ws') + '/_ws?token=' + user.token,
		{
			onOpen: () => console.log('opened'),
			onClose: () => console.log('closed'),
			onMessage: e => {
				console.log(e.data)
				const message = JSON.parse(e.data) as WSMessage

				if (message.type === WSMessageType.RESULT) {
					if (message.data.status === ResultStatus.FAILURE) {
						toast.error(message.data.message)
					}
					return
				}

				const index = conversations.findIndex(c => c.id === message.data.conversationID)

				if (conversation?.id === message.data.conversationID) {
					setMessageHistory(prev => prev.concat(message.data))
				} else {
					conversations[index].unreadCount ??= 0
					conversations[index].unreadCount++
				}

				setConversations(prev => {
					if (index !== -1) {
						const newConversations = [...prev]
						// @ts-expect-error previws.lastMessage does not have recipientUsername
						newConversations[index].lastMessage = message.data
						return newConversations
					}

					return prev
				})
			},
			shouldReconnect: () => true,
			onReconnectStop: () => {
				setUser(null)
				toast.error('Cannot connect to chat server, please login again', { duration: 5000 })
				navigate({ to: '/auth/login', search: { redirect: '/home' } })
			},
			reconnectAttempts: 3,
			reconnectInterval: 2000
		}
	)

	useEffect(() => {
		fetchConversations()
	}, [])

	async function fetchConversations() {
		try {
			const resp = await fetch(API_URL + '/messages/conversations/' + user.id, {
				headers: { 'X-Session': user.token }
			})

			if (resp.ok) {
				const { data } = (await resp.json()) as { data: ConversationPreview[] }
				setConversations(data)
			} else {
				throw new Error(resp.statusText)
			}
		} catch (err: any) {
			console.error(err)
			toast.error('Failed to fetch conversations', { description: err?.message })
		}
	}

	async function handleSend(e: FormEvent<HTMLFormElement>) {
		if (conversation === null) return

		// FIX sometimes it sends an empty message
		try {
			e.preventDefault()
			// @ts-expect-error its fine
			const msg: ConversationMessage = {
				senderID: user.id,
				content: messageValue,
				createdAt: new Date().toUTCString(),
				type: MessageType.NORMAL
			}

			if (!conversation.id && conversation.username) {
				msg.recipientUsername = conversation.username
			} else {
				msg.conversationID = conversation.id
			}

			requestAnimationFrame(() => {
				setTimeout(() => {
					inputRef.current?.focus()
				}, 0)
			})

			sendJsonMessage<ConversationMessage>(msg)
			setMessageHistory(prev => prev.concat(msg))
			setConversations(prev => {
				const index = prev.findIndex(c => c.id === msg.conversationID)
				if (index !== -1) {
					const newConversations = [...prev]
					// @ts-expect-error previws.lastMessage does not have recipientUsername
					newConversations[index].lastMessage = msg
					return newConversations
				}

				return prev
			})
			setMessageValue('')
		} catch (err: any) {
			// TODO handle error
			console.error(err)
			toast.error('Failed to send message', { description: err?.message })
		}
	}

	return (
		<div
			className="w-full h-full"
			onKeyDown={e => {
				if (e.code === 'Escape' && conversation !== null) {
					setConversation(null)
				}
			}}
			tabIndex={0}
		>
			<div className="sm:grid sm:grid-cols-[2fr_5fr] w-full h-full sm:h-96 lg:h-4/5 lg:w-9/12 lg:max-w-5xl absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 rounded-2xl border">
				{readyState === ReadyState.OPEN ? (
					<>
						{/* TODO update last message when a message send or received and resort conversations */}
						{(breakpoint !== 'phone' || !conversation) && <Sidebar conversations={conversationPreviews} />}
						{(breakpoint !== 'phone' || conversation?.id || conversation?.username) && (
							<div ref={autoAnimateRef} className="h-full overflow-hidden">
								{!conversation || conversation.id === 'null' || conversation.id === '""' ? (
									<div className="text-muted-foreground text-lg text-center pt-[calc(35%)]">
										Select a chat to start messaging
									</div>
								) : (
									<>
										<PersonInfo />
										<ChatContainer />
										<MessageBox onSubmit={handleSend} value={messageValue} setValue={setMessageValue} ref={inputRef} />
									</>
								)}
							</div>
						)}
					</>
				) : (
					// TODO i think there should be a better way to center text
					<div className="w-full h-full flex items-center justify-center">
						<div className="text-center">
							<h1 className="text-4xl font-bold">
								{readyState === ReadyState.CONNECTING ? 'Connecting...' : 'Disconnected, please wait...'}
							</h1>
						</div>
					</div>
				)}
			</div>
		</div>
	)
}
