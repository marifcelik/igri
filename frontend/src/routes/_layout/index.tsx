import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_layout/')({
	component: Index
})

function Index() {
	return (
		<div className="p-2">
			<h1 className="text-4xl text-center my-10">Welcome to the chat app</h1>
			<p>this page gonna be completed soon, i guess.</p>
			<p>just login for now</p>
		</div>
	)
}
