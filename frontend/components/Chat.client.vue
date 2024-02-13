<script setup>
import { useWebSocket } from '@vueuse/core'
import { Send } from 'lucide-vue-next'

const message = ref('')

const config = useRuntimeConfig()

const { status } = useWebSocket(config.public.wsURL, {
  autoReconnect: {
    retries: 3,
    delay: 1000,
    onFailed: () => console.log('ws connection failed')
  }
})
</script>

<template>
  <p>{{ status }}</p>
  <p>{{ $route.fullPath }}</p>
  <div class="flex w-full max-w-sm items-center gap-1.5" v-auto-animate>
    <UiInput placeholder="Message" v-model="message" @keyup.enter="message = ''" />
    <UiButton v-if="message.trim().length > 0" type="submit">
      <Send class="w-5 h-5" />
    </UiButton>
  </div>
</template>