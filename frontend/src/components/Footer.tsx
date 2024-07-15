import { Link, type LinkProps } from '@tanstack/react-router'
import { MountainIcon } from 'lucide-react'

function Footer() {
	const links: { name: string; to: LinkProps['to'] }[] = [
		{ name: 'Home', to: '/' },
		{ name: 'About', to: '/about' },
		{ name: 'Prev', to: '../' }
	]

	return (
		<footer className="bg-gray-900 text-gray-400 py-8">
			<div className="container mx-auto px-4 md:px-6">
				<div className="flex flex-col md:flex-row items-center justify-between">
					<div className="flex items-center gap-2">
						<MountainIcon className="h-6 w-6 text-gray-400" />
						<span className="text-lg font-bold">Blog</span>
					</div>
					<nav className="flex items-center gap-4 mt-4 md:mt-0">
						{links.map((link, i) => (
							<Link key={i} className="hover:text-white" to={link.to}>
								{link.name}
							</Link>
						))}
					</nav>
				</div>
				<p className="text-center mt-6 text-sm">© 2024 Blog. All rights reserved.</p>
			</div>
		</footer>
	)
}

export default Footer
