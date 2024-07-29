import { createRootRoute, Outlet, ScrollRestoration } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import SonnerContainer from '@/components/SonnerContainer'

export const Route = createRootRoute({
	component: () => (
		<>
			<Outlet />
			<SonnerContainer />
			<ScrollRestoration />
			<TanStackRouterDevtools />
		</>
	)
})
