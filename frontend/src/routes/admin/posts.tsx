import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/admin/posts')({
	component: () => <div>Hello /admin/posts!</div>
})
