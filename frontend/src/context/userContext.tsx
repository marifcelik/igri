import { createContext, useState } from 'react'
import { UserFields } from '@/types'
import { useLocalStorage } from '@uidotdev/usehooks'

type UserContextType = {
	user: UserFields
	setUser: (user: UserFields) => void
}

export const UserContext = createContext<UserContextType | undefined>(undefined)

export function UserProvider({ children }: { children: React.ReactNode }) {
	const [id, setId] = useLocalStorage('id', '')
	const [username, setUsername] = useLocalStorage('username', '')
	const [token, setToken] = useLocalStorage('token', '')

	const [user, setUser] = useState<UserFields>({ id, username, token })

	function handleSetUser(user: UserFields) {
		setId(user.id)
		setUsername(user.username)
		setToken(user.token)
		setUser(user)
	}

	return <UserContext.Provider value={{ user, setUser: handleSetUser }}>{children}</UserContext.Provider>
}
