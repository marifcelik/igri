import { useContext, useState } from 'react'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { AtSignIcon, CornerDownLeftIcon, MessageCirclePlusIcon } from 'lucide-react'
import { zodResolver } from '@hookform/resolvers/zod'
import useBreakpoint from '@/hooks/breakpoint'
import { Button } from '../ui/button'
import { Input } from '../ui/input'
import { Popover, PopoverContent, PopoverTrigger } from '../ui/popover'
import { Dialog, DialogContent, DialogDescription, DialogTitle, DialogTrigger } from '../ui/dialog'
import { Form, FormControl, FormField, FormItem, FormMessage } from '../ui/form'
import { ChatContext } from '@/context/chatContext'

export default function NewConversation() {
	const [open, setOpen] = useState(false)

	const breakpoint = useBreakpoint()

	// TODO create a separate component for the Button
	return breakpoint === 'phone' ? (
		<Dialog open={open} onOpenChange={setOpen}>
			<DialogTrigger asChild>
				<Button variant="outline" size="icon" className="rounded-full">
					<MessageCirclePlusIcon className="size-5" />
					<span className="sr-only">New chat</span>
				</Button>
			</DialogTrigger>
			<DialogContent>
				<DialogTitle>Start a conversation with</DialogTitle>
				<DialogDescription asChild>
					<UsernameForm setOpen={setOpen} showHeader={false} />
				</DialogDescription>
			</DialogContent>
		</Dialog>
	) : (
		<Popover open={open} onOpenChange={setOpen}>
			<PopoverTrigger asChild>
				<Button variant="outline" size="icon" className="rounded-full">
					<MessageCirclePlusIcon className="size-5" />
					<span className="sr-only">New chat</span>
				</Button>
			</PopoverTrigger>
			<PopoverContent>
				<UsernameForm setOpen={setOpen} />
			</PopoverContent>
		</Popover>
	)
}

function UsernameForm({
	setOpen,
	showHeader = true
}: { setOpen: React.Dispatch<React.SetStateAction<boolean>>; showHeader?: boolean }) {
	const { setConversation } = useContext(ChatContext)!

	const FormSchema = z.object({
		username: z
			.string()
			.trim()
			.min(1, 'Username is required')
			.refine(v => !v.includes(' '), 'Username cannot contain spaces')
	})

	const form = useForm<z.infer<typeof FormSchema>>({
		resolver: zodResolver(FormSchema),
		defaultValues: { username: '' },
		shouldFocusError: true
	})

	function createNewConversation(data: z.infer<typeof FormSchema>) {
		// TODO check if the username exists
		// when a new conversation created, the conversation should only have the username
		setConversation({
			username: data.username
		})
		setOpen(false)
	}

	return (
		<Form {...form}>
			{showHeader && <h2 className="text-sm font-semibold">Start a conversation with</h2>}
			<form onSubmit={form.handleSubmit(createNewConversation)} className="flex items-start justify-between gap-4 mt-3">
				<FormField
					control={form.control}
					name="username"
					render={({ field }) => (
						<FormItem>
							<FormControl>
								<div className="relative w-full items-center">
									<span className="absolute	start-2 inset-y-0 flex items-center justify-center px-1 z-10">
										<AtSignIcon className="size-4 text-muted-foreground" />
									</span>
									<Input {...field} id="username" placeholder="username" className="pl-9" />
								</div>
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				></FormField>
				<Button type="submit" variant={showHeader ? 'secondary' : 'default'}>
					<CornerDownLeftIcon className="size-4" />
				</Button>
			</form>
		</Form>
	)
}
