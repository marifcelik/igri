import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/auth/')({
	component: AuthIndex
})

function AuthIndex() {
	return <div>Auth index</div>
}
