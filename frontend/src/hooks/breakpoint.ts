import { createBreakpoint } from 'react-use'

export const breakpoints = {
	phone: 0,
	sm: 640,
	md: 768,
	lg: 1024,
	xl: 1280,
	'2xl': 1536
}

const useBreakpoint = createBreakpoint(breakpoints) as () => keyof typeof breakpoints

export default useBreakpoint
