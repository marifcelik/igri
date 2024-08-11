import { PhoneIcon, VideoIcon } from 'lucide-react'
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar'
import { Button } from '../ui/button'
import { Link } from '@tanstack/react-router'
import ThemeButton from '../ThemeButton'

export default function PersonInfo() {
	return (
		<div className="p-3 flex border-b items-center h-16">
			<div className="flex items-center gap-2">
				<Avatar className="border w-10 h-10">
					<AvatarImage src="/placeholder-user.jpg" />
					<AvatarFallback>OM</AvatarFallback>
				</Avatar>
				<div className="grid gap-0.5">
					<p className="text-sm font-medium leading-none">Sofia Davis</p>
					<p className="text-xs text-muted-foreground">Active 2h ago</p>
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
				<Button variant="ghost" asChild>
					<Link to="/login">Login</Link>
				</Button>
				<ThemeButton />
			</div>
		</div>
	)
}
