import { useState } from 'react'
import { Link, type LinkProps } from '@tanstack/react-router'
import { SearchIcon, MountainIcon, MenuIcon } from 'lucide-react'
import { Button } from './ui/button'
import { Input } from './ui/input'
import { Sheet, SheetTrigger, SheetContent } from './ui/sheet'
import ThemeButton from './ThemeButton'

function Header() {
	const links: { name: string; to: LinkProps['to'] }[] = [
		{ name: 'Home', to: '/' },
		{ name: 'About', to: '/about' },
		{ name: 'Auth', to: '/auth' }
	]

	const [open, setOpen] = useState(false)

	return (
		<header className="fixed top-0 left-0 w-full backdrop-blur shadow z-50 bg-gray-100/80 dark:bg-[#21252d]/90">
			<div className="container mx-auto flex items-center justify-between px-4 md:px-6 py-4">
				<Link className="flex items-center gap-2" to="/">
					<MountainIcon className="h-6 w-6 text-gray-900 dark:text-gray-50" />
					<span className="text-lg font-bold">Blog</span>
				</Link>
				<div className="flex items-center gap-0 sm:gap-2 md:gap-4">
					<nav className="hidden md:flex items-center gap-4">
						{links.map((link, i) => (
							<Link
								key={i}
								className="text-gray-600 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-50"
								to={link.to}
							>
								{link.name}
							</Link>
						))}
						<ThemeButton />
					</nav>
					<div className="relative w-full max-w-md md:w-auto">
						<Input
							className="w-36 sm:w-48 md:min-w-64 rounded-md border border-gray-300 bg-white py-2 px-4 text-gray-600 focus:border-gray-500 focus:outline-none dark:border-gray-600 dark:bg-gray-700 dark:text-gray-400 dark:focus:border-gray-500"
							placeholder="Search..."
							type="search"
						/>
						<SearchIcon className="absolute right-6 md:right-3 top-1/2 h-5 w-5 -translate-y-1/2" />
					</div>
					<Sheet open={open} onOpenChange={setOpen}>
						<SheetTrigger asChild>
							<Button className="md:hidden border-none" size="icon" variant="outline">
								<MenuIcon className="h-6 w-6" />
								<span className="sr-only">Toggle navigation menu</span>
							</Button>
						</SheetTrigger>
						<SheetContent side="right">
							<nav className="grid gap-4 py-6">
								{links.map((link, i) => (
									<Link key={i} className="flex items-center gap-2" onClick={() => setOpen(false)} to={link.to}>
										{link.name}
									</Link>
								))}
							</nav>
							<ThemeButton />
						</SheetContent>
					</Sheet>
				</div>
			</div>
		</header>
	)
}

export default Header
