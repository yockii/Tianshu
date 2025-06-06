<template>
  <el-card>
    <div style="display:flex; justify-content:space-between; align-items:center; margin-bottom:16px;">
      <h2>操作日志</h2>
    </div>
    <el-table :data="logs" style="width: 100%">
      <el-table-column prop="id" label="日志ID" width="100" />
      <el-table-column prop="tenantId" label="租户ID" width="120" />
      <el-table-column prop="userId" label="用户ID" width="120" />
      <el-table-column prop="action" label="操作类型" />
      <el-table-column prop="detail" label="详情" />
      <el-table-column prop="createdAt" label="时间" width="180">
        <template #default="{ row }">
          {{ new Date(row.createdAt).toLocaleString() }}
        </template>
      </el-table-column>
    </el-table>
    <div style="text-align: right; margin-top: 16px;">
      <el-pagination
        :current-page="page"
        :page-size="pageSize"
        :total="total"
        layout="prev, pager, next"
        @current-change="loadLogs"
      />
    </div>
  </el-card>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { api } from '@/utils/request'

const logs = ref<any[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)

async function loadLogs(current: number) {
  page.value = current
  try {
    const res = await api.get('/logs', { params: { offset: (current - 1) * pageSize.value, limit: pageSize.value } })
    logs.value = res.list
    total.value = res.total
  } catch {
    ElMessage.error('获取操作日志失败')
  }
}

onMounted(() => loadLogs(page.value))
</script>
