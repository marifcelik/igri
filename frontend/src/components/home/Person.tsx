import { useContext } from 'react'
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar'
import { ChatContext } from '@/context/chatContext'
import type { ChatPerson } from '@/types'

export default function Person({ id, name, username: _username, image, time, message }: ChatPerson) {
	const { setReceiver } = useContext(ChatContext)!

	return (
		<div
			className="flex items-center gap-4 p-2 my-2 rounded-lg hover:bg-muted/50 cursor-pointer"
			onClick={() => setReceiver(id)}
		>
			<Avatar className="border w-10 h-10">
				<AvatarImage src={image} />
				<AvatarFallback>
					{name
						.split(' ')
						.map(n => n[0].toLocaleUpperCase())
						.join('')}
				</AvatarFallback>
			</Avatar>
			<div className="grid gap-0.5">
				<p className="text-sm font-medium leading-none">{name}</p>
				<p className="text-xs text-muted-foreground">
					{message} &middot; {time}
				</p>
			</div>
		</div>
	)
}
