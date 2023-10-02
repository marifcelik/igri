<script setup>
import { useColorMode } from '@vueuse/core'

const { store } = useColorMode({
  disableTransition: false
})
const checked = useState('isDark', () => false)

onMounted(() => {
  checked.value = store.value === 'dark'
})

watch([checked], () => {
  store.value = checked.value ? 'dark' : 'light'
})
</script>

<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>
