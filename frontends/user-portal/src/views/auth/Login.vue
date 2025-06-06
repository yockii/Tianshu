<template>
  <el-main class="auth-main">
    <el-card class="auth-card">
      <h2>登录</h2>
      <el-form :model="form" :rules="rules" ref="formRef" label-position="top" style="margin-top: 24px;">
        <el-form-item label="租户ID（可选）" prop="tenantId">
          <el-input-number v-model="form.tenantId" :min="1" placeholder="若无域名，可输入租户ID" style="width: 100%;" />
        </el-form-item>
        <el-form-item label="域名（若有）" prop="domain">
          <el-input v-model="form.domain" placeholder="请输入租户域名" />
        </el-form-item>
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="form.password" type="password" placeholder="请输入密码" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" style="width: 100%;" :loading="loading" @click="onLogin">登录</el-button>
        </el-form-item>
        <el-form-item>
          <el-link @click="$router.push('/register')">还没租户？注册</el-link>
        </el-form-item>
      </el-form>
    </el-card>
  </el-main>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { api } from '@/utils/request'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const formRef = ref()
const loading = ref(false)
const form = ref({
  tenantId: 0,
  domain: '',
  username: '',
  password: ''
})
const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
}

const onLogin = () => {
  formRef.value.validate(async (valid: boolean) => {
    if (!valid) return
    loading.value = true
    try {
      const { token, user, permissions } = await api.post('/user/login', form.value)
      localStorage.setItem('token', token)
      localStorage.setItem('userInfo', JSON.stringify(user))
      localStorage.setItem('tenantId', user.tenantId)
      // initialize user store
      const userStore = useUserStore()
      userStore.setUser(user)
      userStore.setPermissions(permissions)
      ElMessage.success('登录成功')
      router.push('/dashboard')
    } catch (e: any) {
      ElMessage.error(e.message || '登录失败')
    } finally {
      loading.value = false
    }
  })
}
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
