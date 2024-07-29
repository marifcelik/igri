import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_layout/auth/')({
	component: AuthIndex
})

function AuthIndex() {
	return <div>Auth index</div>
}
