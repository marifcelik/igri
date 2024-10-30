// import { createPortal } from 'react-dom'
import { createFileRoute, Outlet, redirect } from '@tanstack/react-router'
import Header from '@/components/Header'
// import Footer from '@/components/Footer'

export const Route = createFileRoute('/_layout')({
	beforeLoad: () => {
		const token = localStorage.getItem('token')
		if (token && token !== '""') throw redirect({ to: '/home' })
	},
	component: () => (
		<>
			<Header />
			<main className="container mx-auto pt-24 md:pt-28 pb-12 md:pb-16">
				<Outlet />
			</main>
			{/* createPortal(<Footer />, document.body) */}
		</>
	)
})
