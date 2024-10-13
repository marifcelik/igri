import { createFileRoute, Outlet, redirect } from '@tanstack/react-router'
import Sidebar from '@/components/admin/Sidebar'

export const Route = createFileRoute('/admin')({
	beforeLoad: async ({ location }) => {
		const token = localStorage.getItem('admin-token')
		if (!token)
			throw redirect({
				to: '/auth/login',
				search: { redirect: location.href }
			})
	},
	component: () => (
		<>
			<Sidebar />
			<main className="continer flex absolute top-0 right-0 w-full md:w-[calc(100%-18rem)]">
				<Outlet />
			</main>
		</>
	)
})
