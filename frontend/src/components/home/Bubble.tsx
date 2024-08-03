export default function Bubble({
	position,
	children
}: {
	position: 'left' | 'right'
	children: React.ReactNode
}) {
	return (
		<div
			className={`flex flex-col w-max max-w-[65%] rounded-full px-4 py-2 my-4 text-sm  ${
				position === 'right' ? 'ml-auto bg-primary text-primary-foreground' : 'bg-muted'
			}`}
		>
			{children}
		</div>
	)
}
