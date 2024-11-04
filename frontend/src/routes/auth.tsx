import { createFileRoute, Link, Outlet, redirect } from '@tanstack/react-router'
import { ArrowLeftIcon } from 'lucide-react'
import ThemeButton from '@/components/ThemeButton'

export const Route = createFileRoute('/auth')({
	component: AuthLayout,
	beforeLoad: () => {
		const token = localStorage.getItem('token')
		if (token && !['', 'null', '"null"', 'undefined', '"undefined"', '""'].includes(token))
			throw redirect({ to: '/home' })
	}
})

function AuthLayout() {
	return (
		<>
			<div className="absolute top-10 left-10 md:right-32">
				<ThemeButton />
			</div>
			<div className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 ">
				<Link
					to="/"
					className="inline-flex items-center justify-center w-full gap-2 mb-10 text-blue-600 hover:text-blue-800 dark:text-blue-400 dark:hover:text-blue-500"
				>
					<ArrowLeftIcon className="h-5 w-5" />
					Return to homepage
				</Link>
				<Outlet />
			</div>
		</>
	)
}
