import { useAutoAnimate } from '@formkit/auto-animate/react'
import ConversationItem from './ConversationItem'
import NewConversation from './NewConversation'
import ThemeButton from '../ThemeButton'
import { Input } from '../ui/input'
import type { ConversationPreview } from '@/types'
import { LogOutIcon } from 'lucide-react'
import { Button } from '../ui/button'

export default function Sidebar({ conversations }: { conversations: ConversationPreview[] }) {
	const [autoAnimateRef] = useAutoAnimate()

	function handleLogout() {
		localStorage.setItem('token', '""')
		location.reload()
	}

	return (
		<div className="bg-muted/20 p-3 border-r overflow-hidden h-full">
			<div className="h-24">
				<div className="flex items-center justify-between space-x-4">
					<div className="font-medium text-sm ml-4">Chats</div>
					<div className="flex items-center justify-center gap-3">
						<Button className="rounded-full" variant="destructive" size="icon" onClick={handleLogout}>
							<LogOutIcon className="size-5 rotate-180" />
						</Button>
						<NewConversation />
						<ThemeButton />
					</div>
				</div>
				<div className="py-4">
					<Input placeholder="Search" className="h-8" />
				</div>
			</div>
			<div ref={autoAnimateRef} className="h-[calc(100%-6rem)] overflow-y-auto">
				{conversations.map((conversation, index) => (
					<ConversationItem key={index} item={conversation} />
				))}
			</div>
		</div>
	)
}
