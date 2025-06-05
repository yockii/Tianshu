<template>
  <el-main class="auth-main">
    <el-card class="auth-card">
      <h2>租户注册</h2>
      <el-form :model="form" :rules="rules" ref="formRef" label-position="top" style="margin-top: 24px;">
        <el-form-item label="租户名称" prop="tenantName">
          <el-input v-model="form.tenantName" placeholder="请输入租户名称" />
        </el-form-item>
        <el-form-item label="管理员用户名" prop="adminUsername">
          <el-input v-model="form.adminUsername" placeholder="请输入管理员用户名" />
        </el-form-item>
        <el-form-item label="管理员邮箱" prop="adminEmail">
          <el-input v-model="form.adminEmail" placeholder="请输入管理员邮箱" />
        </el-form-item>
        <el-form-item label="管理员手机号" prop="adminPhone">
          <el-input v-model="form.adminPhone" placeholder="请输入手机号（可选）" />
        </el-form-item>
        <el-form-item label="管理员密码" prop="adminPassword">
          <el-input v-model="form.adminPassword" type="password" placeholder="请输入管理员密码" />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input v-model="form.confirmPassword" type="password" placeholder="请再次输入密码" />
        </el-form-item>
        <el-form-item label="租户域名" prop="domain">
          <el-input v-model="form.domain" placeholder="可选，租户专属域名" />
        </el-form-item>
        <el-form-item label="Logo" prop="logo">
          <el-input v-model="form.logo" placeholder="可选，Logo地址" />
        </el-form-item>
        <el-form-item label="主题色" prop="theme">
          <el-input v-model="form.theme" placeholder="可选，主题色" />
        </el-form-item>
        <el-form-item label="欢迎语" prop="welcomeText">
          <el-input v-model="form.welcomeText" placeholder="可选，欢迎语" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" style="width: 100%;" :loading="loading" @click="onRegister">注册</el-button>
        </el-form-item>
        <el-form-item>
          <el-link @click="$router.push('/login')">已有账号？登录</el-link>
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

const router = useRouter()
const formRef = ref()
const loading = ref(false)
const form = ref({
  tenantName: '',
  adminUsername: '',
  adminEmail: '',
  adminPhone: '',
  adminPassword: '',
  confirmPassword: '',
  domain: '',
  logo: '',
  theme: '',
  welcomeText: ''
})
const rules = {
  tenantName: [{ required: true, message: '请输入租户名称', trigger: 'blur' }],
  adminUsername: [{ required: true, message: '请输入管理员用户名', trigger: 'blur' }],
  adminEmail: [{ required: true, message: '请输入管理员邮箱', trigger: 'blur' }],
  adminPassword: [{ required: true, message: '请输入管理员密码', trigger: 'blur' }],
  confirmPassword: [{ required: true, message: '请再次输入密码', trigger: 'blur' },
    { validator: (rule: any, value: string, callback: any) => {
      if (value !== form.value.adminPassword) {
        callback(new Error('两次密码输入不一致'))
      } else {
        callback()
      }
    }, trigger: 'blur' }
  ]
}

const onRegister = () => {
  formRef.value.validate(async (valid: boolean) => {
    if (!valid) return
    loading.value = true
    try {
      const res = await api.post('/user/register', form.value)
      if (res.tenant && res.admin) {
        ElMessage.success('注册成功，请登录')
        router.push('/login')
      } else {
        ElMessage.error(res.message || '注册失败')
      }
    } catch (e: any) {
      ElMessage.error(e.message || '注册失败')
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
  width: 400px;
}
</style>
