import { useState } from 'react'
import { Link } from '@tanstack/react-router'
import { MountainIcon, MenuIcon } from 'lucide-react'
import { Button } from './ui/button'
import { Sheet, SheetTrigger, SheetContent } from './ui/sheet'
import ThemeButton from './ThemeButton'

function Header() {
	// const links: { name: string; to: LinkProps['to'] }[] = [{ name: 'Login', to: '/auth/login' }]

	const [open, setOpen] = useState(false)

	return (
		<header className="fixed top-0 left-0 w-full z-30 sm:px-28  bg-transparent">
			<div className="container mx-auto flex items-center justify-between px-4 md:px-6 py-4">
				<Link className="flex items-center gap-2" to="/">
					<MountainIcon className="h-6 w-6 text-gray-900 dark:text-gray-50" />
					<span className="text-lg font-bold">İgri</span>
				</Link>
				<div className="flex items-center gap-0 sm:gap-2 md:gap-4">
					<nav className="hidden md:flex items-center gap-4">
						{/* {links.map((link, i) => (
							<Link key={i} className={buttonVariants({ variant: 'ghost' })} to={link.to}>
								{link.name}
							</Link>
						))} */}
						<ThemeButton />
					</nav>
					<Sheet open={open} onOpenChange={setOpen}>
						<SheetTrigger asChild>
							<Button className="md:hidden border-none" size="icon" variant="outline">
								<MenuIcon className="h-6 w-6" />
								<span className="sr-only">Toggle navigation menu</span>
							</Button>
						</SheetTrigger>
						<SheetContent side="right">
							<nav className="grid gap-4 py-6">
								{/* {links.map((link, i) => (
									<Link key={i} className="flex items-center gap-2" onClick={() => setOpen(false)} to={link.to}>
										{link.name}
									</Link>
								))} */}
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
