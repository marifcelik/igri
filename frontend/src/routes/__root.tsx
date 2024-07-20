import { createPortal } from 'react-dom'
import { createRootRoute, Outlet, ScrollRestoration } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import Header from '@/components/Header'
import Footer from '@/components/Footer'

export const Route = createRootRoute({
	component: () => (
		<>
			<Header />
			<main className="container mx-auto pt-24 md:pt-28 pb-12 md:pb-16">
				<Outlet />
			</main>
			{createPortal(<Footer />, document.body)}
			<ScrollRestoration />
			<TanStackRouterDevtools />
		</>
	)
})
