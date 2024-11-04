export default function Bubble({
	position,
	children
}: {
	position: 'left' | 'right'
	children: React.ReactNode
}) {
	return (
		<div
			className={`flex flex-col w-max h-max max-w-[55%] rounded-2xl px-4 py-2 my-2 text-sm ${
				position === 'right' ? 'ml-auto bg-primary text-primary-foreground' : 'bg-neutral-200 dark:bg-zinc-700'
			}`}
		>
			{children}
		</div>
	)
}
