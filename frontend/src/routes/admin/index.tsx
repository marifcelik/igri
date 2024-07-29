import { createFileRoute } from '@tanstack/react-router'
import { toast } from 'sonner'
import { Button } from '@/components/ui/button'

export const Route = createFileRoute('/admin/')({
	component: Index
})

function Index() {
	return (
		<div className="w-full md:w-3/4 mx-auto mt-28 px-3 md:px-0 grid grid-cols-4 gap-2">
			<Button onClick={() => toast.success('deneemeeeee')}>success</Button>
			<Button variant="destructive" onClick={() => toast.error('deneemeeeee')}>
				error
			</Button>
			<Button onClick={() => toast.info('deneemeeeee')}>info</Button>
			<Button onClick={() => toast.warning('deneemeeeee')}>warning</Button>
		</div>
	)
}
