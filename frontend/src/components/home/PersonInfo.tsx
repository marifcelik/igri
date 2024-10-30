import { useContext } from 'react'
import { ArrowLeftIcon, PhoneIcon, VideoIcon } from 'lucide-react'
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar'
import { Button } from '../ui/button'
import { ChatContext } from '@/context/chatContext'

export default function PersonInfo() {
	const { conversation, setConversation } = useContext(ChatContext)!

	return (
		<div className="p-3 flex border-b items-center h-16">
			<div className="flex items-center gap-2">
				<Button className="sm:hidden sm:opacity-0" variant="ghost" size="icon" onClick={() => setConversation(null)}>
					<ArrowLeftIcon className="size-4" />
				</Button>
				<Avatar className="border w-10 h-10">
					<AvatarImage src={conversation?.image} />
					<AvatarFallback>
						{(conversation?.name ?? conversation!.username)
							.split(' ')
							.map(n => n[0].toLocaleUpperCase())
							.join('')
							.slice(0, 2)}
					</AvatarFallback>
				</Avatar>
				<div className="grid gap-0.5">
					<p className="text-sm font-medium leading-none">{conversation?.name ?? conversation?.username}</p>
					{/* TODO implement last active time */}
					<p className="text-xs text-muted-foreground">{conversation?.username} - TODO</p>
				</div>
			</div>
			<div className="flex items-center gap-1 ml-auto">
				<Button variant="ghost" size="icon">
					<span className="sr-only">Call</span>
					<PhoneIcon className="h-4 w-4" />
				</Button>
				<Button variant="ghost" size="icon">
					<span className="sr-only">Video call</span>
					<VideoIcon className="h-4 w-4" />
				</Button>
			</div>
		</div>
	)
}
