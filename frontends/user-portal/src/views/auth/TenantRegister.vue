<template>
  <el-main class="auth-main">
    <el-card class="auth-card">
      <h2>租户注册</h2>
      <el-form label-position="top" style="margin-top: 24px;">
        <el-form-item label="租户名称">
          <el-input placeholder="请输入租户名称" />
        </el-form-item>
        <el-form-item label="管理员用户名">
          <el-input placeholder="请输入管理员用户名" />
        </el-form-item>
        <el-form-item label="管理员邮箱">
          <el-input placeholder="请输入管理员邮箱" />
        </el-form-item>
        <el-form-item label="管理员密码">
          <el-input type="password" placeholder="请输入管理员密码" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" style="width: 100%;">注册租户</el-button>
        </el-form-item>
        <el-form-item>
          <el-link @click="$router.push('/login')">已有账号？登录</el-link>
        </el-form-item>
      </el-form>
    </el-card>
  </el-main>
</template>

<script setup lang="ts">
</script>

<style scoped>
.auth-main {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 80vh;
}
.auth-card {
  width: 360px;
}
</style>
