import { Toaster } from 'sonner'
import { useTheme } from '@/components/ThemeProvider'

export default function SonnerContainer() {
	const { theme } = useTheme()
	return <Toaster richColors position="top-right" duration={3000} visibleToasts={4} theme={theme} />
}
