import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/admin/messages')({
	component: () => <div>Hello /_admin/messages!</div>
})
