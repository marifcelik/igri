import { createLazyFileRoute, Link, useNavigate } from '@tanstack/react-router'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { toast } from 'sonner'
import { ArrowLeftIcon } from 'lucide-react'
import ThemeButton from '@/components/ThemeButton'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card'
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form'
import { API_URL } from '@/lib/config'

export const Route = createLazyFileRoute('/auth/login')({
	component: Login
})

const LoginFormSchema = z.object({
	username: z.string().min(1, 'Username is required'),
	password: z.string().min(6, 'Password must be at least 6 characters')
})

function Login() {
	const form = useForm<z.infer<typeof LoginFormSchema>>({
		resolver: zodResolver(LoginFormSchema),
		defaultValues: {
			username: '',
			password: ''
		}
	})

	const navigate = useNavigate()

	async function handleLogin(data: z.infer<typeof LoginFormSchema>) {
		const resp = await fetch(API_URL + '/auth/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(data)
		})

		if (resp.ok) {
			const { token } = await resp.json()
			toast.success('Welcome back')
			localStorage.setItem('token', token)
			navigate({ to: '/' })
		} else {
			const text = await resp.text()
			toast.error(text)
			console.log(resp)
		}
	}

	return (
		<>
			<div className="absolute top-30 right-10 md:right-32">
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
													<FormLabel htmlFor="username">Name</FormLabel>
													<FormControl>
														<Input {...field} id="username" placeholder="Enter your username" />
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
													<FormLabel htmlFor="password">Email</FormLabel>
													<FormControl>
														<Input {...field} id="password" placeholder="Enter your password" type="password" />
													</FormControl>
													<FormMessage />
												</FormItem>
											)}
										/>
									</div>
								</div>
								<Button type="submit" className="mt-3">
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
