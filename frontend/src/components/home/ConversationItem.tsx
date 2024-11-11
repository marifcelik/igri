import { useContext } from 'react'
import { formatDistanceToNow } from 'date-fns'
import { ChatContext } from '@/context/chatContext'
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar'
import { Badge } from '../ui/badge'
import type { ConversationPreview } from '@/types'

export default function ConversationItem({ item }: { item: ConversationPreview }) {
	const { conversation, setConversation } = useContext(ChatContext)!

	return (
		<div
			className={`flex items-center gap-4 p-2 my-2 rounded-lg cursor-pointer ${conversation?.id === item.id ? 'bg-neutral-200 dark:bg-zinc-700' : 'hover:bg-muted/50'}`}
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
					{item.lastMessage.content.length > 34
						? item.lastMessage.content.length < 39
							? item.lastMessage.content
							: item.lastMessage.content.slice(0, 34) + '...'
						: item.lastMessage.content}
					{' '}&middot;{' '}
					{formatDistanceToNow(new Date(item.lastMessage.createdAt), { addSuffix: true, includeSeconds: true })}
				</p>
			</div>
			{item.unreadCount !== undefined && item.unreadCount > 0 && <Badge>{item.unreadCount}</Badge>}
		</div>
	)
}
