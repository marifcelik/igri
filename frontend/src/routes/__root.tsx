import { useState } from 'react'
import { useLocalStorage } from 'react-use'
import { createRootRoute, Outlet, ScrollRestoration } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import SonnerContainer from '@/components/SonnerContainer'
import { ChatContext } from '@/context/chatContext'
import { UserContext } from '@/context/userContext'
import type { ConversationPreview, UserFields, ConversationMessage } from '@/types'

export const Route = createRootRoute({
	component: __Root
})

function __Root() {
	const [id, setId] = useLocalStorage<string>('id', '')
	const [username, setUsername] = useLocalStorage<string>('username', '')
	const [token, setToken] = useLocalStorage<string>('token', '')

	const [recipientID, setRecipientID] = useState<string | null>(null)
	const [conversation, setConversation] = useState<ConversationPreview | null>(null)
	// @ts-expect-error react-use's useLocalStorage returns undefined even though we've set a default value
	const [user, setUser] = useState<UserFields>({ id, username, token })
	const [messageHistory, setMessageHistory] = useState<ConversationMessage[]>([])

	function handleSetUser(user: UserFields | null) {
		if (!user) {
			setId('')
			setUsername('')
			setToken('')
			setUser({ id: '', username: '', token: '' })
			return
		}

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
					// @ts-expect-error its fine
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
