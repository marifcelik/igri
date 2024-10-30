import { useContext, useState } from 'react'
import { createLazyFileRoute, Link } from '@tanstack/react-router'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { toast } from 'sonner'
import { EyeIcon, EyeOffIcon, Loader2Icon } from 'lucide-react'
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
	status: string
}

const LoginFormSchema = z.object({
	username: z.string().trim().min(1, 'Username is required'),
	password: z.string().min(6, 'Password must be at least 6 characters')
})

export const Route = createLazyFileRoute('/auth/login')({
	component: Login
})

function Login() {
	// const router = useRouter()
	const { redirect } = Route.useSearch<{ redirect?: string }>()
	// const navigate = useNavigate()

	const [loading, setLoading] = useState(false)
	const [showPassword, setShowPassword] = useState(false)

	// TODO animate the password visibility icon
	// const [autoAnimateRef] = useAutoAnimate<HTMLDivElement>()

	const { setUser } = useContext(UserContext)!

	const form = useForm<z.infer<typeof LoginFormSchema>>({
		resolver: zodResolver(LoginFormSchema),
		defaultValues: {
			username: '',
			password: ''
		},
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
				setUser({
					id: data.id,
					username: data.username,
					token: token
				})
				toast.success('Welcome back ' + data.username)

				// FIX after login, conversations don't load, so, for now, im using hard relocation
				if (redirect !== undefined && redirect !== '') window.location.href = redirect
				else window.location.href = '/'
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
			toast.error('An error occurred', { description: err.message, duration: 5000 })
		} finally {
			setLoading(false)
		}
	}

	return (
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
												<div className="relative w-full max-w-sm items-center">
													<span
														className="absolute	end-2 inset-y-0 flex items-center justify-center px-2 cursor-pointer z-10"
														onClick={() => setShowPassword(!showPassword)}
													>
														{showPassword ? <EyeIcon className="size-5" /> : <EyeOffIcon className="size-5" />}
													</span>
													<Input
														{...field}
														id="password"
														placeholder="Enter your password"
														type={showPassword ? 'text' : 'password'}
														disabled={loading}
														className="pr-10"
													/>
												</div>
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
				<div className="text-sm mx-auto mt-3">
					Don't have an account?
					<Link
						to="/auth/register"
						search={redirect !== undefined ? { redirect } : undefined}
						className="ml-2 text-center text-blue-600 hover:text-blue-800 dark:text-blue-400 dark:hover:text-blue-500"
					>
						Register
					</Link>
				</div>
			</CardContent>
		</Card>
	)
}
