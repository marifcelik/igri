<script setup lang="ts">
const checked = useState('isDark', () => false)

definePageMeta({
  layout: 'home'
})
</script>

<template>
  <ClientOnly>
    <main>
      index page
      <UiLabel for="color-mode">Switch Theme</UiLabel>
      <UiSwitch id="color-mode" v-model:checked="checked" /> <br />
      <span>color mode is {{ checked ? 'dark' : 'light' }}</span>
    </main>
  </ClientOnly>
</template>
