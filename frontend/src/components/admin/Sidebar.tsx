import { useState } from 'react'
import { Link, LinkProps } from '@tanstack/react-router'
import { HomeIcon, MailIcon, MenuIcon, ListIcon, MountainIcon, LogOutIcon, type LucideIcon } from 'lucide-react'
import { Sheet, SheetTrigger, SheetContent } from '../ui/sheet'
import ThemeButton from '../ThemeButton'
import { Button } from '../ui/button'

const links: { name: string; icon: LucideIcon; to: LinkProps['to'] }[] = [
	{ name: 'Homepage', icon: HomeIcon, to: '/admin' },
	{ name: 'Posts', icon: ListIcon, to: '/admin/posts' },
	{ name: 'Messages', icon: MailIcon, to: '/admin/messages' }
]

function Sidebar() {
	function handleLogout() {
		localStorage.removeItem('token')
		window.location.reload()
	}

	const [open, setOpen] = useState(false)

	return (
		<aside id="sidebar" className="fixed top-0 left-0 w-16 md:w-72 z-50">
			<nav className="hidden md:block bg-background border-r-2 h-screen p-5 pt-8 relative">
				<div className="h-12 flex items-center justify-between px-4">
					<MountainIcon className="w-8 h-8" />
					<ThemeButton />
				</div>
				<ul className="pt-6">
					{links.map((item, i) => (
						<li
							key={i}
							className="rounded-lg mt-2 text-xl cursor-pointer hover:bg-primary/30 has-[.active]:bg-primary/70 has-[.active]:text-white dark:has-[.active]:bg-primary/50 has-[.active]:hover:bg-primary/50 duration-200"
						>
							<Link to={item.to} className="w-full h-full flex items-center rounded-lg p-4 gap-x-4">
								<item.icon />
								<span className="origin-left duration-200">{item.name}</span>
							</Link>
						</li>
					))}
				</ul>
				<Button
					variant="destructive"
					className="absolute bottom-0 left-0 w-full p-4 py-7 gap-x-4 text-xl flex items-center rounded-none"
					onClick={handleLogout}
				>
					<LogOutIcon className="rotate-180" />
					<span>Log out</span>
				</Button>
			</nav>
			<Sheet open={open} onOpenChange={setOpen}>
				<SheetTrigger asChild>
					<Button className="md:hidden border-none absolute top-5 left-5" size="icon" variant="outline">
						<MenuIcon className="h-6 w-6" />
						<span className="sr-only">Toggle navigation menu</span>
					</Button>
				</SheetTrigger>
				<SheetContent side="left">
					<nav className="mb-3">
						<ul>
							{links.map((item, i) => (
								<li
									key={i}
									className="rounded-lg mt-2 cursor-pointer dark:hover:bg-stone-100/10 hover:bg-neutral-200/50"
								>
									<Link
										to={item.to}
										className="w-full h-full flex items-center rounded-lg p-4 gap-x-4"
										onClick={() => setOpen(false)}
									>
										<item.icon />
										<span className="origin-left duration-200">{item.name}</span>
									</Link>
								</li>
							))}
						</ul>
					</nav>
					<ThemeButton />
				</SheetContent>
			</Sheet>
		</aside>
	)
}

export default Sidebar
