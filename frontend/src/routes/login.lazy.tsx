import { useContext, useState } from 'react'
import { createLazyFileRoute, Link, useNavigate, useRouter } from '@tanstack/react-router'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { toast } from 'sonner'
import { ArrowLeftIcon, Loader2Icon } from 'lucide-react'
import ThemeButton from '@/components/ThemeButton'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card'
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form'
import { UserContext } from '@/context/userContext'
import { API_URL } from '@/lib/config'

type LoginData = {
	data: {
		id: string
		name: string
		username: string
		createdAt: string
	}
}

const LoginFormSchema = z.object({
	username: z.string().trim().min(1, 'Username is required'),
	password: z.string().min(6, 'Password must be at least 6 characters')
})

export const Route = createLazyFileRoute('/login')({
	component: Login
})

function Login() {
	const router = useRouter()
	const { redirect } = Route.useSearch<{ redirect?: string }>()
	const navigate = useNavigate()

	const [loading, setLoading] = useState(false)

	const { setUser } = useContext(UserContext)!

	const form = useForm<z.infer<typeof LoginFormSchema>>({
		resolver: zodResolver(LoginFormSchema),
		defaultValues: {
			username: '',
			password: ''
		},
		// FIX doesn't focus the form
		shouldFocusError: true
	})

	async function handleLogin(data: z.infer<typeof LoginFormSchema>) {
		setLoading(true)
		// TODO clear the code
		try {
			const resp = await fetch(API_URL + '/auth/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(data)
			})

			if (resp.ok) {
				const token = resp.headers.get('x-session')
				if (token === null) throw new Error('Token not found in response headers')

				const { data } = (await resp.json()) as LoginData
				localStorage.setItem('to', data.username === '"marifcelik"' ? '"tıpıt"' : '"marifcelik"')
				setUser({
					id: data.id,
					username: data.username,
					token: token
				})
				toast.success('Welcome back ' + data.username)

				if (redirect !== undefined && redirect !== '') router.history.push(redirect)
				else navigate({ to: '/' })
			} else {
				const text = await resp.text()
				if (resp.headers.get('Content-Type')?.includes('application/json')) {
					const err = JSON.parse(text) as { status: string; message: string; data: Record<string, string[]> | string }
					toast.error(err.message, {
						description:
							err.data instanceof Object
								? Object.values(err.data).map((t, i) => <p key={i}>&middot; {t}</p>)
								: err.data,
						duration: 5000
					})
				} else {
					toast.error('An error occurred', { description: text, duration: 5000 })
				}
			}
		} catch (err) {
			console.error('handle login error', err)
			// @ts-expect-error err is unknown
			toast.error('An error occurred', { description: err.message as string, duration: 5000 })
		} finally {
			setLoading(false)
		}
	}

	return (
		<>
			<div className="absolute top-10 right-10 md:right-32">
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
				<Card className="w-80 scale-110">
					<CardHeader className="space-y-1">
						<CardTitle className="text-2xl">Login</CardTitle>
					</CardHeader>
					<CardContent className="grid gap-4">
						<Form {...form}>
							<form onSubmit={form.handleSubmit(handleLogin)} className="grid gap-4 my-auto">
								<div className="grid gap-4">
									<div className="space-y-2">
										<FormField
											control={form.control}
											name="username"
											render={({ field }) => (
												<FormItem>
													<FormLabel htmlFor="username">Username</FormLabel>
													<FormControl>
														<Input {...field} id="username" placeholder="Enter your username" disabled={loading} />
													</FormControl>
													<FormMessage />
												</FormItem>
											)}
										/>
									</div>
									<div className="space-y-2">
										<FormField
											control={form.control}
											name="password"
											render={({ field }) => (
												<FormItem>
													<FormLabel htmlFor="password">Password</FormLabel>
													<FormControl>
														<Input
															{...field}
															id="password"
															placeholder="Enter your password"
															type="password"
															disabled={loading}
														/>
													</FormControl>
													<FormMessage />
												</FormItem>
											)}
										/>
									</div>
								</div>
								<Button type="submit" className="mt-3" disabled={loading}>
									{loading && <Loader2Icon className="mr-2 size-4 animate-spin" />}
									Login
								</Button>
							</form>
						</Form>
					</CardContent>
				</Card>
			</div>
		</>
	)
}
