// filepath: src/stores/user.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { api } from '@/utils/request'
import type { ApiResponse } from '@/types'

interface UserInfo {
  id: number
  tenantId: number
  username: string
  email: string
  phone?: string
  isSuperAdmin: boolean
  status: number
}

export const useUserStore = defineStore('user', () => {
  const userInfo = ref<UserInfo | null>(null)
  const permissions = ref<string[]>([])

  function setUser(info: UserInfo) {
    userInfo.value = info
  }
  function setPermissions(codes: string[]) {
    permissions.value = codes
  }

  async function fetchPermissions() {
    try {
      const codes = await api.get<string[]>('/relation/user-permissions')
      permissions.value = codes
    } catch (e) {
      console.error('Failed to fetch permissions', e)
    }
  }

  function hasPermission(code: string): boolean {
    if (userInfo.value?.isSuperAdmin) return true
    return permissions.value.includes(code)
  }

  return { userInfo, permissions, setUser, setPermissions, fetchPermissions, hasPermission }
})
