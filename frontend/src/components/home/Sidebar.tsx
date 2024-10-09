import { PlusIcon } from 'lucide-react'
import Person from './Person'
import { Button } from '../ui/button'
import { Input } from '../ui/input'
import type { ChatPerson } from '@/types'

export default function Sidebar() {
	const persons: ChatPerson[] = [
		{
			id: crypto.randomUUID(),
			name: 'Tıpıt Çelik',
			username: 'tıpıt',
			message: "hey what's going on?",
			time: '2h',
			image: '/placeholder-user.jpg'
		},
		{
			id: crypto.randomUUID(),
			name: 'Maho adam',
			username: 'marifcelik',
			message: 'Just finished a great book! 📚',
			time: '45m',
			image: '/placeholder-user.jpg'
		},
		{
			id: crypto.randomUUID(),
			name: 'Tadam adam',
			username: 'tadamadam',
			message: 'Excited for the weekend!',
			time: '1h',
			image: '/placeholder-user.jpg'
		},
		{
			id: crypto.randomUUID(),
			name: 'valorant fatihi',
			username: 'valorantfatihi',
			message: "Who's up for a movie night?",
			time: '3h',
			image: '/placeholder-user.jpg'
		},
		{
			id: crypto.randomUUID(),
			name: 'Walter White',
			username: 'walterwhite',
			message: 'Morning coffee is the best! ☕',
			time: '30m',
			image: '/placeholder-user.jpg'
		},
		{
			id: crypto.randomUUID(),
			name: 'Anakin Skywalker',
			username: 'anakinskywalker',
			message: 'I am the chosen one!',
			time: '1d',
			image: '/placeholder-user.jpg'
		},
		{
			id: crypto.randomUUID(),
			name: 'Finn Mertens',
			username: 'finnmertens',
			message: 'Adventure time!',
			time: '1d',
			image: '/placeholder-user.jpg'
		},
		{
			id: crypto.randomUUID(),
			name: 'Bruce Wayne',
			username: 'brucewayne',
			message: 'I am not Batman',
			time: 'just now',
			image: '/placeholder-user.jpg'
		},
		{
			id: crypto.randomUUID(),
			name: 'Cooper',
			username: 'cooper',
			message: 'We are not alone',
			time: '??',
			image: '/placeholder-user.jpg'
		},
		{
			id: crypto.randomUUID(),
			name: 'Rorschach',
			username: 'rorschach',
			message: 'The end is nigh',
			time: '1w',
			image: '/placeholder-user.jpg'
		},
		{
			id: crypto.randomUUID(),
			name: 'Miles Morales',
			username: 'milesmorales',
			message: 'Hey, I might be late today',
			time: '1h',
			image: '/placeholder-user.jpg'
		},
		{
			id: crypto.randomUUID(),
			name: 'James Howlett',
			username: 'jameshowlett',
			message: 'fuck off',
			time: 'just now',
			image: '/placeholder-user.jpg'
		}
	]

	return (
		<div className="bg-muted/20 p-3 border-r overflow-hidden h-full">
			<div className="h-24">
				<div className="flex items-center justify-between space-x-4">
					<div className="font-medium text-sm">Chats</div>
					<Button variant="ghost" size="icon" className="rounded-full w-8 h-8">
						<PlusIcon className="h-4 w-4" />
						<span className="sr-only">New chat</span>
					</Button>
				</div>
				<div className="py-4">
					<Input placeholder="Search" className="h-8" />
				</div>
			</div>
			<div className="h-[calc(100%-6rem)] overflow-y-scroll">
				{persons.map((person, index) => (
					<Person key={index} {...person} />
				))}
			</div>
		</div>
	)
}
