<script setup lang="ts">
const users = ['furkan', 'arif', 'kadir', 'michella']
</script>

<template>
  <aside class="h-screen">
    <div class="w-full h-28 border flex items-center justify-center">
      <Profile />
    </div>
    <UiScrollArea class="ml-0 mr-auto h-[calc(100%-7rem)] border border-t-0">
      <div v-for="(user, i) in users" :key="i" class="h-24 w-full hover:cursor-pointer">
        <div class="flex h-full w-full items-center justify-between space-x-4 hover:bg-slate-100 dark:hover:bg-sky-950 transition-colors">
          <div class="flex items-center space-x-4 ml-6">
            <UiAvatar class="h-16 w-16">
              <UiAvatarImage src="" />
              <UiAvatarFallback>FN</UiAvatarFallback>
            </UiAvatar>
            <div>
              <p class="text-sm font-medium leading-none">
                {{ user }}
              </p>
              <p class="text-sm text-muted-foreground">
                @{{ user }}
              </p>
            </div>
          </div>
        </div>
        <UiSeparator />
      </div>
    </UiScrollArea>
  </aside>
</template>
