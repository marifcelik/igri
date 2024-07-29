import { useState } from 'react'
import { createFileRoute } from '@tanstack/react-router'
import { toast } from 'sonner'
import { Button } from '@/components/ui/button'

export const Route = createFileRoute('/_layout/auth/register')({
	component: Register
})

function Register() {
	const fns: Array<keyof Pick<typeof toast, 'info' | 'success' | 'error' | 'warning'>> = [
		'info',
		'success',
		'error',
		'warning'
	]
	const [index, setIndex] = useState<number>(0)

	function handleClick() {
		toast[fns[index]]('deneme')
		setIndex(index === fns.length - 1 ? 0 : index + 1)
	}

	return (
		<div>
			<h2>register</h2>
			<Button onClick={handleClick}>sonner deneme</Button>
		</div>
	)
}
