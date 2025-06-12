import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { type User } from '@/types/user'
import pilotBridge from '@/dji/pilot-bridge'

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(null)
  const isLoggedIn = computed(() => user.value !== null && token.value !== null && token.value !== '')
  const username = computed(() => user.value ? user.value.username : '')
  const userId = computed(() => user.value ? user.value.id : '')

  const setUser = (newUser: User | null) => {
    user.value = newUser
  }
  const setToken = (newToken: string | null) => {
    token.value = newToken
    pilotBridge.setToken(newToken? newToken : undefined)
  }
  const logout = () => {
    user.value = null
    token.value = null
  }
  return { user, token, isLoggedIn, username, userId, setUser, setToken, logout }
}, {
  persist: true,
})
