<template>
  <el-container class="dashboard-layout">
    <el-aside width="220px">
      <el-menu :default-active="$route.name" router>
        <el-menu-item index="Dashboard" @click="$router.push('/dashboard')">工作台</el-menu-item>
        <el-menu-item index="Devices" @click="$router.push('/dashboard/devices')">设备管理</el-menu-item>
        <el-menu-item index="Missions" @click="$router.push('/dashboard/missions')">任务管理</el-menu-item>
        <el-menu-item index="Users" @click="$router.push('/dashboard/users')">用户管理</el-menu-item>
        <el-menu-item index="TenantSettings" @click="$router.push('/dashboard/tenant')">租户设置</el-menu-item>
        <el-menu-item index="Profile" @click="$router.push('/dashboard/profile')">个人中心</el-menu-item>
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
function logout() {
  localStorage.removeItem('token')
  localStorage.removeItem('userInfo')
  window.location.href = '/login'
}
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
