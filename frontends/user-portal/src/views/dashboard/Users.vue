<template>
  <el-card>
    <h2>用户管理</h2>
    <el-table :data="users">
      <el-table-column prop="id" label="用户ID" />
      <el-table-column prop="username" label="用户名" />
      <el-table-column prop="email" label="邮箱" />
      <el-table-column prop="createdAt" label="注册时间" />
    </el-table>
    <el-pagination
      background
      layout="total, prev, pager, next" 
      :total="total"
      :page-size="limit"
      @current-change="onPageChange"
    />
  </el-card>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { api } from '@/utils/request'

const users = ref<any[]>([])
const total = ref(0)
const limit = ref(10)
const offset = ref(0)

const fetchUsers = () => {
  api.get('/user/list', { offset: offset.value, limit: limit.value })
    .then(res => {
      users.value = res.list
      total.value = res.total
    })
    .catch(() => ElMessage.error('获取用户列表失败'))
}

onMounted(fetchUsers)

const onPageChange = (page: number) => {
  offset.value = (page - 1) * limit.value
  fetchUsers()
}
</script>
