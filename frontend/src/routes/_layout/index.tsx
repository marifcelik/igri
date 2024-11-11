import { createFileRoute } from '@tanstack/react-router'
import { useTheme } from '@/components/ThemeProvider'
import Globe from '@/components/ui/globe'
import AnimatedGradientText from '@/components/ui/animated-gradient-text'
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
				<AnimatedGradientText text="İgri" />
				<div className="flex flex-col sm:flex-row mt-5 items-center justify-center">
					<h2 className="sm:w-1/3 text-muted-foreground font-semibold m-5">İgri is a simple chat application</h2>
					<ShimmerButton onClick={() => navigate({ to: '/home' })} className="shadow-2xl z-20">
						<span className="whitespace-pre-wrap text-center text-sm font-medium leading-none tracking-tight text-white dark:from-white dark:to-slate-900/10 lg:text-lg">
							Login Now
						</span>
					</ShimmerButton>
				</div>
				<DotPattern className="[mask-image:radial-gradient(300px_circle_at_center,white,transparent)] -z-0" />
			</div>
			{/* FIX when page is firstly loaded, globe is moving on mouse hover */}
			<Globe dark={theme === 'dark' ? 0 : 1} />
		</div>
	)
}
