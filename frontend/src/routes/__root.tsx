import { useState } from 'react'
import { useLocalStorage } from '@uidotdev/usehooks'
import { createRootRoute, Outlet, ScrollRestoration } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import SonnerContainer from '@/components/SonnerContainer'
import { ChatContext } from '@/context/chatContext'
import { UserContext } from '@/context/userContext'
import type { UserFields, WSMessage } from '@/types'

export const Route = createRootRoute({
	component: () => {
		const [id, setId] = useLocalStorage<string>('id')
		const [username, setUsername] = useLocalStorage<string>('username')
		const [token, setToken] = useLocalStorage<string>('token')

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
				<ChatContext.Provider value={{ messageHistory, setMessageHistory }}>
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
