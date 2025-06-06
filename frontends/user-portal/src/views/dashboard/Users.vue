<template>
  <el-card>
    <div style="display:flex; justify-content:space-between; align-items:center; margin-bottom:16px;">
      <h2>用户管理</h2>
      <el-button type="primary" @click="openDialog('create')">新增用户</el-button>
    </div>
    <el-table :data="users">
      <el-table-column prop="id" label="用户ID" />
      <el-table-column prop="username" label="用户名" />
      <el-table-column prop="email" label="邮箱" />
      <el-table-column prop="createdAt" label="注册时间" />
      <el-table-column label="操作">
        <template #default="{ row }">
          <el-button type="text" size="small" @click="openDialog('edit', row)">编辑</el-button>
          <el-button type="text" size="small" @click="deleteUser(row)">删除</el-button>
          <!-- 不允许为超级管理员分配角色 -->
          <el-button v-if="!row.isSuperAdmin" type="text" size="small" @click="openAssignRoles(row)">分配角色</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      background
      layout="total, prev, pager, next" 
      :total="total"
      :page-size="limit"
      @current-change="onPageChange"
    />
    <!-- 用户新增/编辑对话框 -->
    <el-dialog v-model="showDialog" :title="isEdit ? '编辑用户' : '新增用户'">
      <el-form :model="userForm" label-width="80px">
        <el-form-item label="用户名">
          <el-input v-model="userForm.username" />
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="userForm.email" />
        </el-form-item>
        <el-form-item label="电话">
          <el-input v-model="userForm.phone" />
        </el-form-item>
        <el-form-item label="密码" v-if="!isEdit">
          <el-input type="password" v-model="userForm.password" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="userForm.status" placeholder="请选择状态">
            <el-option :label="'正常'" :value="1" />
            <el-option :label="'禁用'" :value="0" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showDialog = false">取消</el-button>
        <el-button type="primary" @click="submitForm">确定</el-button>
      </template>
    </el-dialog>
    <!-- Role assignment dialog -->
    <el-dialog v-model="assignDialogVisible" title="分配角色">
      <!-- 超级管理员不允许修改角色 -->
      <div v-if="selectedUser && selectedUser.isSuperAdmin" style="color: #f56c6c; margin-bottom: 16px;">超级管理员角色不可修改</div>
      <el-checkbox-group v-model="selectedRoleIds">
        <el-checkbox v-for="role in rolesList" :key="role.id" :label="role.id" :disabled="selectedUser && selectedUser.isSuperAdmin">{{ role.name }}</el-checkbox>
      </el-checkbox-group>
      <template #footer>
        <el-button @click="assignDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitRoles" :disabled="selectedUser && selectedUser.isSuperAdmin">确定</el-button>
      </template>
    </el-dialog>
  </el-card>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { api } from '@/utils/request'

// User list and pagination
const users = ref<any[]>([])
const total = ref(0)
const limit = ref(10)
const offset = ref(0)

// User create/edit dialog
const showDialog = ref(false)
const isEdit = ref(false)
const userForm = reactive({ id: 0, username: '', email: '', phone: '', password: '', status: 1 })

// Role assignment dialog
const rolesList = ref<any[]>([])
const assignDialogVisible = ref(false)
const selectedUser = ref<any>(null)
const selectedRoleIds = ref<number[]>([])

// Fetch users
async function fetchUsers() {
  try {
    const res = await api.get<{ list: any[]; total: number }>('/user/list', { offset: offset.value, limit: limit.value })
    users.value = res.list
    total.value = res.total
  } catch {
    ElMessage.error('获取用户列表失败')
  }
}

// Fetch roles for assignment
async function fetchRoles() {
  try {
    const res = await api.get<{ list: any[]; total: number }>('/role/list')
    rolesList.value = res.list
  } catch {
    ElMessage.error('获取角色列表失败')
  }
}

onMounted(async () => {
  await fetchUsers()
  await fetchRoles()
})

function onPageChange(page: number) {
  offset.value = (page - 1) * limit.value
  fetchUsers()
}

// Open create/edit user dialog
function openDialog(mode: 'create' | 'edit', row?: any) {
  isEdit.value = (mode === 'edit')
  if (isEdit.value && row) {
    Object.assign(userForm, { id: row.id, username: row.username, email: row.email, phone: row.phone, status: row.status, password: '' })
  } else {
    Object.assign(userForm, { id: 0, username: '', email: '', phone: '', status: 1, password: '' })
  }
  showDialog.value = true
}

// Submit create/edit user
async function submitForm() {
  try {
    if (isEdit.value) {
      await api.put(`/user/${userForm.id}`, { ...userForm, passwordHash: userForm.password })
    } else {
      await api.post('/user/create', { ...userForm, password: userForm.password })
    }
    ElMessage.success(isEdit.value ? '更新成功' : '创建成功')
    showDialog.value = false
    fetchUsers()
  } catch {
    ElMessage.error(isEdit.value ? '更新失败' : '创建失败')
  }
}

// Delete user
function deleteUser(row: any) {
  ElMessageBox.confirm('确定要删除该用户？', '提示', { type: 'warning' })
    .then(async () => {
      try {
        await api.delete(`/user/${row.id}`)
        ElMessage.success('删除成功')
        fetchUsers()
      } catch {
        ElMessage.error('删除失败')
      }
    }).catch(() => {})
}

// Open role assignment dialog
async function openAssignRoles(row: any) {
  selectedUser.value = row
  try {
    const assigned: any[] = await api.get<any[]>('/relation/user-roles', { userId: row.id })
    selectedRoleIds.value = assigned.filter(r => r.userId === row.id).map(r => r.roleId)
    assignDialogVisible.value = true
  } catch {
    ElMessage.error('获取用户角色失败')
  }
}

// Submit role assignments
async function submitRoles() {
  if (!selectedUser.value) return
  const userId = selectedUser.value.id
  try {
    // fetch current assignments for this user
    const assigned: any[] = await api.get<any[]>('/relation/user-roles', { userId })
    const current = assigned.filter(r => r.userId === userId).map(r => r.roleId)
    const toAdd = selectedRoleIds.value.filter(id => !current.includes(id))
    const toRemove = current.filter(id => !selectedRoleIds.value.includes(id))
    for (const rid of toAdd) {
      await api.post('/relation/user-role', { userId, roleId: rid })
    }
    for (const rid of toRemove) {
      await api.delete('/relation/user-role', { userId, roleId: rid })
    }
    ElMessage.success('角色分配更新成功')
    assignDialogVisible.value = false
  } catch {
    ElMessage.error('角色分配更新失败')
  }
}
</script>
