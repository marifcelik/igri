<script setup>
import { GithubIcon } from 'lucide-vue-next'
</script>

<template>
  <UiTabs defaultValue="login" class="w-96 mx-auto">
    <UiTabsList class="grid w-full grid-cols-2">
      <UiTabsTrigger value="login">
        Login
      </UiTabsTrigger>
      <UiTabsTrigger value="register">
        Register
      </UiTabsTrigger>
    </UiTabsList>
    <UiTabsContent value="login">
      <UiCard>
        <UiCardHeader class="space-y-1">
          <UiCardTitle class="text-2xl">
            Login
          </UiCardTitle>
          <UiCardDescription>
            Enter your email below
          </UiCardDescription>
        </UiCardHeader>
        <UiCardContent class="grid gap-4">
          <div class="grid grid-cols-2 gap-6">
            <UiButton variant="outline">
              <GithubIcon class="mr-2 h-4 w-4" />
              Github
            </UiButton>
            <UiButton variant="outline">
              <svg role="img" viewBox="0 0 24 24" class="mr-2 h-4 w-4">
                <path fill="currentColor"
                  d="M12.48 10.92v3.28h7.84c-.24 1.84-.853 3.187-1.787 4.133-1.147 1.147-2.933 2.4-6.053 2.4-4.827 0-8.6-3.893-8.6-8.72s3.773-8.72 8.6-8.72c2.6 0 4.507 1.027 5.907 2.347l2.307-2.307C18.747 1.44 16.133 0 12.48 0 5.867 0 .307 5.387.307 12s5.56 12 12.173 12c3.573 0 6.267-1.173 8.373-3.36 2.16-2.16 2.84-5.213 2.84-7.667 0-.76-.053-1.467-.173-2.053H12.48z" />
              </svg>
              Google
            </UiButton>
          </div>
          <div class="relative">
            <div class="absolute inset-0 flex items-center">
              <span class="w-full border-t" />
            </div>
            <div class="relative flex justify-center text-xs uppercase">
              <span class="bg-background px-2 text-muted-foreground">
                Or continue with
              </span>
            </div>
          </div>
          <div class="grid gap-2">
            <UiLabel for="email">Email</UiLabel>
            <UiInput id="email" type="email" placeholder="m@example.com" />
          </div>
          <div class="grid gap-2">
            <UiLabel for="password">Password</UiLabel>
            <UiInput id="password" type="password" />
          </div>
        </UiCardContent>
        <UiCardFooter>
          <UiButton class="w-full">
            Create account
          </UiButton>
        </UiCardFooter>
      </UiCard>
    </UiTabsContent>
    <LazyUiTabsContent value="register">
      <UiCard>
        <UiCardHeader class="space-y-1">
          <UiCardTitle class="text-2xl">
            Create an account
          </UiCardTitle>
          <UiCardDescription>
            Enter your email below to create your account
          </UiCardDescription>
        </UiCardHeader>
        <UiCardContent class="grid gap-4">
          <div class="grid grid-cols-2 gap-6">
            <UiButton variant="outline">
              <GithubIcon class="mr-2 h-4 w-4" />
              Github
            </UiButton>
            <UiButton variant="outline">
              <svg role="img" viewBox="0 0 24 24" class="mr-2 h-4 w-4">
                <path fill="currentColor"
                  d="M12.48 10.92v3.28h7.84c-.24 1.84-.853 3.187-1.787 4.133-1.147 1.147-2.933 2.4-6.053 2.4-4.827 0-8.6-3.893-8.6-8.72s3.773-8.72 8.6-8.72c2.6 0 4.507 1.027 5.907 2.347l2.307-2.307C18.747 1.44 16.133 0 12.48 0 5.867 0 .307 5.387.307 12s5.56 12 12.173 12c3.573 0 6.267-1.173 8.373-3.36 2.16-2.16 2.84-5.213 2.84-7.667 0-.76-.053-1.467-.173-2.053H12.48z" />
              </svg>
              Google
            </UiButton>
          </div>
          <div class="relative">
            <div class="absolute inset-0 flex items-center">
              <span class="w-full border-t" />
            </div>
            <div class="relative flex justify-center text-xs uppercase">
              <span class="bg-background px-2 text-muted-foreground">
                Or continue with
              </span>
            </div>
          </div>
          <div class="grid gap-2">
            <UiLabel for="email">Email</UiLabel>
            <UiInput id="email" type="email" placeholder="m@example.com" />
          </div>
          <div class="grid gap-2">
            <UiLabel for="password">Password</UiLabel>
            <UiInput id="password" type="password" />
          </div>
        </UiCardContent>
        <UiCardFooter>
          <UiButton class="w-full">
            Create account
          </UiButton>
        </UiCardFooter>
      </UiCard>
    </LazyUiTabsContent>
  </UiTabs>
</template>