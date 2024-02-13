// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  experimental: {
    typedPages: true
  },
  modules: [
    '@nuxtjs/tailwindcss',
    '@vueuse/nuxt',
    '@vueuse/motion/nuxt',
    '@formkit/auto-animate/nuxt',
    '@vee-validate/nuxt'
  ],
  components: [
    {
      path: '~/components/ui',
      extensions: ['.vue'],
      prefix: 'Ui'
    },
    '~/components'
  ],
  tailwindcss: {
    cssPath: '~/assets/css/main.css',
  },
  runtimeConfig: {
    public: {
      wsURL: ''
    }
  }
})