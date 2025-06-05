<template>
  <el-card>
    <h2>租户设置</h2>
    <el-form :model="form" label-position="top" ref="formRef">
      <el-form-item label="租户名称">
        <el-input v-model="form.name" placeholder="请输入租户名称" />
      </el-form-item>
      <el-form-item label="Logo 地址">
        <el-input v-model="form.logo" placeholder="请输入 Logo 地址" />
      </el-form-item>
      <el-form-item label="主题色">
        <el-color-picker v-model="form.theme" />
      </el-form-item>
      <el-form-item label="欢迎语">
        <el-input v-model="form.welcomeText" placeholder="请输入欢迎语" />
      </el-form-item>
      <!-- 定制化字段 -->
      <el-form-item label="站点名称">
        <el-input v-model="form.siteName" placeholder="请输入站点名称" />
      </el-form-item>
      <el-form-item label="Favicon 地址">
        <el-input v-model="form.favicon" placeholder="请输入 Favicon 地址" />
      </el-form-item>
      <!-- <el-form-item label="额外配置 (JSON)">
        <el-input type="textarea" :rows="4" v-model="form.extraConfig" placeholder="输入 JSON 格式的额外配置" />
      </el-form-item> -->
      <el-form-item>
        <el-button type="primary" :loading="loading" @click="onSave">保存设置</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { api } from '@/utils/request'

// 扩展接口包含定制化字段
interface TenantForm {
  id: string
  name: string
  logo?: string
  theme?: string
  welcomeText?: string
  siteName?: string
  favicon?: string
  extraConfig?: string
}
// 初始化表单模型
const form = ref<TenantForm>({ id: '', name: '', logo: '', theme: '', welcomeText: '', siteName: '', favicon: '', extraConfig: '' })
const formRef = ref()
const loading = ref(false)

// Fetch tenant settings on mount
onMounted(async () => {
  const tenantId = localStorage.getItem('tenantId')
  try {
    const res = await api.get<TenantForm>('/tenant/profile')
    // 支持直接返回对象或包装形式
    const data = (res as any).data !== undefined ? (res as any).data : res
    form.value = { ...data }
  } catch (e: any) {
    ElMessage.error('获取租户信息失败')
  }
})

// Save tenant settings
const onSave = async () => {
  loading.value = true
  try {
    // 发送更新请求，包含定制化字段
    await api.put('/tenant/profile', form.value)
    ElMessage.success('保存成功')
  } catch (e: any) {
    ElMessage.error('保存失败')
  } finally {
    loading.value = false
  }
}
</script>
