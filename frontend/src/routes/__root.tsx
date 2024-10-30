import { useState } from 'react'
import { createRootRoute, Outlet, ScrollRestoration } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import { useLocalStorage } from 'react-use'
import SonnerContainer from '@/components/SonnerContainer'
import { ChatContext } from '@/context/chatContext'
import { UserContext } from '@/context/userContext'
import type { ConversationPreview, UserFields, WSMessage } from '@/types'

export const Route = createRootRoute({
	component: () => {
		const [id, setId] = useLocalStorage<string>('id', '')
		const [username, setUsername] = useLocalStorage<string>('username', '')
		const [token, setToken] = useLocalStorage<string>('token', '')

		const [recipientID, setRecipientID] = useState<string | null>(null)
		const [conversation, setConversation] = useState<ConversationPreview | null>(null)
		// @ts-expect-error
		const [user, setUser] = useState<UserFields>({ id, username, token })
		const [messageHistory, setMessageHistory] = useState<WSMessage[]>([])

		function handleSetUser(user: UserFields) {
			setId(user.id)
			setUsername(user.username)
			setToken(user.token)
			setUser(user)
		}

		return (
			<>
				<ChatContext.Provider
					value={{
						messageHistory,
						setMessageHistory,
						recipientID,
						setRecipientID,
						conversation,
						setConversation
					}}
				>
					<UserContext.Provider value={{ user, setUser: handleSetUser }}>
						<Outlet />
					</UserContext.Provider>
				</ChatContext.Provider>

				<SonnerContainer />
				<ScrollRestoration />
				{import.meta.env.DEV && <TanStackRouterDevtools />}
			</>
		)
	}
})
