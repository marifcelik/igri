<script setup lang="ts">
import {
  Cloud,
  CreditCard,
  Github,
  Keyboard,
  LifeBuoy,
  LogOut,
  Mail,
  MessageSquare,
  Plus,
  PlusCircle,
  Settings,
  User,
  UserPlus,
  Users
} from 'lucide-vue-next'
import { DropdownMenuPortal } from './ui/dropdown-menu'

const dropdownMenuItems: DropdownItems[][] = [
  [
    { icon: User, label: 'Profile', shortcut: '⇧⌘P' },
    { icon: CreditCard, label: 'Billing', shortcut: '⌘B' },
    { icon: Settings, label: 'Settings', shortcut: '⌘S' },
    { icon: Keyboard, label: 'Keyboard shortcuts', shortcut: '⌘K' }
  ],
  [
    { icon: Users, label: 'Team' },
    {
      icon: UserPlus,
      label: 'Invite users',
      subMenu: [
        { icon: Mail, label: 'Email' },
        { icon: MessageSquare, label: 'Message' },
        { icon: PlusCircle, label: 'More...' }
      ]
    },
    { icon: Plus, label: 'New Team', shortcut: '⌘+T' }
  ],
  [
    { icon: Github, label: 'GitHub' },
    { icon: LifeBuoy, label: 'Support' },
    { icon: Cloud, label: 'API', disabled: true }
  ],
  [{ icon: LogOut, label: 'Log out', shortcut: '⇧⌘Q' }]
]
</script>

<template>
  <UiDropdownMenu>
    <UiDropdownMenuTrigger as-child>
      <UiAvatar class="h-16 w-16 cursor-pointer">
        <UiAvatarImage src="finn.jpg" />
        <UiAvatarFallback>FN</UiAvatarFallback>
      </UiAvatar>
    </UiDropdownMenuTrigger>
    <UiDropdownMenuContent class="w-56" side="right" align="start">
      <UiDropdownMenuLabel>My Account</UiDropdownMenuLabel>
      <UiDropdownMenuSeparator />
      <template v-for="(group, i) in dropdownMenuItems" :key="i">
        <UiDropdownMenuGroup>
          <template v-for="item, j in group" :key="j">
            <UiDropdownMenuItem v-if="!item.subMenu" :class="/* @ts-ignore */{ 'focus:bg-destructive': i === dropdownMenuItems.length - 1 }">
              <component :is="item.icon" class="mr-2 h-4 w-4" />
              <span>{{ item.label }}</span>
              <UiDropdownMenuShortcut v-if="item.shortcut">{{
                item.shortcut
              }}</UiDropdownMenuShortcut>
            </UiDropdownMenuItem>

            <UiDropdownMenuSub v-else>
              <UiDropdownMenuSubTrigger>
                <component :is="item.icon" class="mr-2 h-4 w-4" />
                <span>{{ item.label }}</span>
              </UiDropdownMenuSubTrigger>
              <DropdownMenuPortal>
                <UiDropdownMenuSubContent>
                  <UiDropdownMenuItem v-for="subItem in item.subMenu">
                    <component :is="subItem.icon" class="mr-2 h-4 w-4" />
                    <span>{{ subItem.label }}</span>
                  </UiDropdownMenuItem>
                </UiDropdownMenuSubContent>
              </DropdownMenuPortal>
            </UiDropdownMenuSub>
          </template>
        </UiDropdownMenuGroup>
        <UiDropdownMenuSeparator v-if="i < dropdownMenuItems.length - 1" />
      </template>
    </UiDropdownMenuContent>
  </UiDropdownMenu>
</template>
