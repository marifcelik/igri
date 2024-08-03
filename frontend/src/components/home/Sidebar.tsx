import { PlusIcon } from 'lucide-react'
import Person from './Person'
import { Button } from '../ui/button'
import { Input } from '../ui/input'

export default function Sidebar() {
	const persons = [
		{
			name: 'Tıpıt Çelik',
			message: "hey what's going on?",
			time: '2h',
			image: '/placeholder-user.jpg'
		},
		{
			name: 'Maho adam',
			message: 'Just finished a great book! 📚',
			time: '45m',
			image: '/placeholder-user.jpg'
		},
		{
			name: 'Tadam adam',
			message: 'Excited for the weekend!',
			time: '1h',
			image: '/placeholder-user.jpg'
		},
		{
			name: 'valorant fatihi',
			message: "Who's up for a movie night?",
			time: '3h',
			image: '/placeholder-user.jpg'
		},
		{
			name: 'Walter White',
			message: 'Morning coffee is the best! ☕',
			time: '30m',
			image: '/placeholder-user.jpg'
		},
		{
			name: 'Anakin Skywalker',
			message: 'I am the chosen one!',
			time: '1d',
			image: '/placeholder-user.jpg'
		},
		{
			name: 'Finn Mertens',
			message: 'Adventure time!',
			time: '1d',
			image: '/placeholder-user.jpg'
		},
		{
			name: 'Bruce Wayne',
			message: 'I am not Batman',
			time: 'just now',
			image: '/placeholder-user.jpg'
		},
		{
			name: 'Cooper',
			message: 'We are not alone',
			time: '??',
			image: '/placeholder-user.jpg'
		},
		{
			name: 'Rorschach',
			message: 'The end is nigh',
			time: '1w',
			image: '/placeholder-user.jpg'
		},
		{
			name: 'Miles Morales',
			message: 'Hey, I might be late today',
			time: '1h',
			image: '/placeholder-user.jpg'
		},
		{
			name: 'James Howlett',
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
