import { createLazyFileRoute } from '@tanstack/react-router'

export const Route = createLazyFileRoute('/auth/register')({
	component: Register
})

function Register() {
	return <div>register</div>
}
