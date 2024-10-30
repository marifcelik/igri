import { useContext } from 'react'
import { formatDistanceToNow } from 'date-fns'
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar'
import { ChatContext } from '@/context/chatContext'
import type { ConversationPreview } from '@/types'

export default function ConversationItem({ item }: { item: ConversationPreview }) {
	const { conversation, setConversation } = useContext(ChatContext)!

	return (
		<div
			className={`flex items-center gap-4 p-2 my-2 rounded-lg hover:bg-muted/50 cursor-pointer ${conversation?.id === item.id ? 'bg-muted/50' : ''}`}
			onClick={() => setConversation(item)}
		>
			<Avatar className="border w-10 h-10">
				<AvatarImage src={item.image} />
				<AvatarFallback>
					{item.name
						.split(' ')
						.map(n => n[0].toLocaleUpperCase())
						.join('')
						.slice(0, 2)}
				</AvatarFallback>
			</Avatar>
			<div className="grid gap-0.5">
				<p className="text-sm font-medium leading-none">{item.name}</p>
				<p className="text-xs text-muted-foreground">
					{item.lastMessage.content} &middot;{' '}
					{formatDistanceToNow(new Date(item.lastMessage.createdAt), { addSuffix: true })}
				</p>
			</div>
		</div>
	)
}
