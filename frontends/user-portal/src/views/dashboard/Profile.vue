<template>
  <el-card>
    <h2>个人中心</h2>
    <el-form label-position="top" :model="form" ref="formRef">
      <el-form-item label="用户名">
        <el-input v-model="form.username" disabled />
      </el-form-item>
      <el-form-item label="邮箱">
        <el-input v-model="form.email" disabled />
      </el-form-item>
      <el-form-item label="修改密码">
        <el-input v-model="form.password" type="password" placeholder="新密码" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" :loading="loading" @click="onSave">保存</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { api } from '@/utils/request'

interface ProfileForm {
  id: string
  tenantId: string
  username: string
  email: string
  phone?: string
  password?: string
}
const form = ref<ProfileForm>({ id: '', tenantId: '', username: '', email: '', phone: '', password: '' })
const loading = ref(false)
const formRef = ref()

// Fetch profile on mount
onMounted(async () => {
  const userInfo = JSON.parse(localStorage.getItem('userInfo') || '{}')
  const userId = userInfo.id
  try {
    const res = await api.get<ProfileForm>('/user/profile', { id: userId })
    // 支持带 code 包装或原始对象返回
    const payload = (res as any).data !== undefined ? (res as any).data : res
    form.value = { ...payload, password: '' }
  } catch (e: any) {
    ElMessage.error('获取用户信息失败')
  }
})

// Save password update
const onSave = () => {
  loading.value = true
  const payload: any = {
    id: form.value.id,
    tenantId: form.value.tenantId,
    username: form.value.username,
    email: form.value.email,
    phone: form.value.phone
  }
  if (form.value.password) {
    payload.passwordHash = form.value.password
  }
  api.put('/user/profile', payload)
    .then(() => {
      ElMessage.success('更新成功')
      form.value.password = ''
    })
    .catch(() => {
      ElMessage.error('更新失败')
    })
    .finally(() => {
      loading.value = false
    })
}
</script>
