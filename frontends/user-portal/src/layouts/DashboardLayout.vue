<template>
  <el-container class="dashboard-layout">
    <el-aside width="220px">
      <el-menu :default-active="$route.name" router>
        <el-menu-item v-for="item in menuItems" :index="item.name" :key="item.name" @click="$router.push(item.route)">{{ item.label }}</el-menu-item>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header height="60px" class="dashboard-header">
        <span>天枢无人机智能管控平台</span>
        <el-button type="text" @click="logout">退出登录</el-button>
      </el-header>
      <el-main>
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useUserStore } from '@/stores/user'

function logout() {
  localStorage.removeItem('token')
  localStorage.removeItem('userInfo')
  window.location.href = '/login'
}
const userStore = useUserStore()
const menuItems = computed(() => {
  const items = [
    { name: 'Dashboard', label: '工作台', route: '/dashboard' },
    { name: 'Devices', label: '设备管理', route: '/dashboard/devices' },
    { name: 'Missions', label: '任务管理', route: '/dashboard/missions' }
  ]
  if (userStore.hasPermission('user:list')) {
    items.push({ name: 'Users', label: '用户管理', route: '/dashboard/users' })
  }
  if (userStore.hasPermission('tenant:update')) {
    items.push({ name: 'TenantSettings', label: '租户设置', route: '/dashboard/tenant' })
  }
  if (userStore.hasPermission('role:list')) {
    items.push({ name: 'RoleSettings', label: '角色管理', route: '/dashboard/roles' })
  }
  if (userStore.hasPermission('logs:list')) {
    items.push({ name: 'Logs', label: '操作日志', route: '/dashboard/logs' })
  }
  return items
})
</script>

<style scoped>
.dashboard-layout {
  min-height: 100vh;
}
.dashboard-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #fff;
  box-shadow: 0 2px 8px #f0f1f2;
  padding: 0 24px;
}
</style>
