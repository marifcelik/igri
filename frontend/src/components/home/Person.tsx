import { Link } from '@tanstack/react-router'
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar'

export default function Person({
	name,
	image,
	time,
	message
}: { name: string; image: string; time: string; message: string }) {
	return (
		<Link href="#" className="flex items-center gap-4 p-2 rounded-lg hover:bg-muted/50">
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
		</Link>
	)
}
