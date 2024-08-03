import { createFileRoute, Link, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/_layout/auth')({
	component: AuthLayout
})

function AuthLayout() {
	return (
		<div className="p-2">
			auth layout <br />
			<div className="p-5 border-2 w-56">
				<Link to="/login" className="[&.active]:font-bold">
					login
				</Link>
				<br />
				<Link to="/auth/register" className="[&.active]:font-bold">
					register
				</Link>
			</div>
			<Outlet />
		</div>
	)
}
