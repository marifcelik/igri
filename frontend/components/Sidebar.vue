<script setup lang="ts">
const users = ['furkan', 'arif', 'kadir', 'michella']
</script>

<template>
  <aside class="h-screen">
    <div class="flex h-28 w-full items-center justify-center border">
      <Profile />
    </div>
    <UiScrollArea class="ml-0 mr-auto h-[calc(100%-7rem)] border border-t-0">
      <div
        v-for="(user, i) in users"
        :key="i"
        class="h-24 w-full hover:cursor-pointer !relative"
      >
        <div class="flex h-full w-full items-center justify-between space-x-4 transition-colors hover:bg-slate-300 hover:bg-opacity-30 dark:hover:bg-opacity-10">
          <div class="ml-6 flex items-center space-x-4">
            <UiAvatar class="h-16 w-16">
              <UiAvatarImage src="" />
              <UiAvatarFallback>FN</UiAvatarFallback>
            </UiAvatar>
            <div>
              <p class="text-sm font-medium leading-none">{{ user }}</p>
              <p class="text-sm text-muted-foreground">@{{ user }}</p>
            </div>
          </div>
        </div>
        <UiSeparator />
      </div>
    </UiScrollArea>
  </aside>
</template>
