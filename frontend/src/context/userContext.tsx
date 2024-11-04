import { createContext } from 'react'
import type { UserFields } from '@/types'

type UserContextType = {
	user: UserFields
	setUser: (user: UserFields | null) => void
}

export const UserContext = createContext<UserContextType | undefined>(undefined)
