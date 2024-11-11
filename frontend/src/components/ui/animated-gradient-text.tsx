import { cn } from '@/lib/utils'

/**
 * AnimatedGradientText component
 * @param {number} duration - The duration of the animation in milliseconds
 * @param {string[]} colors - The colors of the gradient, could be any valid CSS color
 */
export default function AnimatedGradientText({
	className = '',
	text,
	duration = 18000,
	colors = []
}: {
	className?: string
	text: string
	duration?: number
	colors?: string[]
}) {
	if (!colors || colors.length < 1) {
		for (let i = 0; i <= 8; i++) {
			colors.push(`hsl(${i * 45}, 100%, 50%)`)
		}
	}

	const style = {
		background: `linear-gradient(90deg, ${colors.join(', ')}) 0 0 / var(--bg-size) 100%`,
		'--duration': `${duration}ms`
	}

	return (
		<h1 className={cn(className + 'rainbow-text')} style={style}>
			{text}
		</h1>
	)
}
