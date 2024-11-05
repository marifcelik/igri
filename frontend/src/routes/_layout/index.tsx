import { createFileRoute } from '@tanstack/react-router'
import { useTheme } from '@/components/ThemeProvider'
import Globe from '@/components/ui/globe'
import AnimatedGradientText from '@/components/ui/animated-gradient-text/'
import DotPattern from '@/components/ui/dot-pattern'
import ShimmerButton from '@/components/ui/shimmer-button'

export const Route = createFileRoute('/_layout/')({
	component: Index
})

function Index() {
	const { theme } = useTheme()

	const navigate = Route.useNavigate()

	return (
		<div className="grid grid-cols-1 sm:grid-cols-2">
			<div className="relative flex h-[300px] sm:h-full w-full flex-col items-center justify-center overflow-hidden bg-background">
				<AnimatedGradientText text="İgri ebi" />
				<div className="flex mt-5 items-center justify-center">
					<ShimmerButton onClick={() => navigate({ to: '/home' })} className="shadow-2xl z-20">
						<span className="whitespace-pre-wrap text-center text-sm font-medium leading-none tracking-tight text-white dark:from-white dark:to-slate-900/10 lg:text-lg">
							tee ışrap çeye
						</span>
					</ShimmerButton>
				</div>
				<DotPattern className="[mask-image:radial-gradient(300px_circle_at_center,white,transparent)] -z-0" />
			</div>
			<Globe dark={theme === 'dark' ? 0 : 1} />
		</div>
	)
}
