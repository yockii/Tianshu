<template>
  <el-card>
    <div style="display:flex; justify-content:space-between; align-items:center; margin-bottom:16px;">
      <h2>角色管理</h2>
      <el-button type="primary" @click="openDialog('create')">新增角色</el-button>
    </div>
    <el-table :data="roles">
      <el-table-column prop="id" label="角色ID" />
      <el-table-column prop="name" label="角色名称" />
      <el-table-column prop="description" label="描述" />
      <el-table-column label="操作">
        <template #default="{ row }">
          <el-button type="text" size="small" @click="openPermDialog(row)">分配权限</el-button>
          <el-button type="text" size="small" @click="openDialog('edit', row)">编辑</el-button>
          <el-button type="text" size="small" @click="deleteRole(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog v-model="showDialog" :title="dialogTitle">
      <el-form :model="roleForm" label-width="80px">
        <el-form-item label="名称">
          <el-input v-model="roleForm.name" />
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="roleForm.description" />
        </el-form-item>
        <el-form-item label="默认角色">
          <el-checkbox v-model="roleForm.isDefault">是</el-checkbox>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showDialog = false">取消</el-button>
        <el-button type="primary" @click="submitForm">确定</el-button>
      </template>
    </el-dialog>
    <!-- 权限分配对话框 -->
    <el-dialog v-model="showPermDialog" title="分配权限" width="40%" top="5vh">
      <el-tree
        ref="permTreeRef"
        :data="treeData"
        show-checkbox
        node-key="id"
        default-expand-all
        :props="{ children: 'children', label: 'label' }"
        :default-checked-keys="selectedPermIds"
        @check-change="onTreeCheckChange"
        style="height:calc(100vh - 200px); overflow:auto;"
      />
      <template #footer>
        <el-button @click="showPermDialog = false">取消</el-button>
        <el-button type="primary" @click="submitPerms">保存</el-button>
      </template>
    </el-dialog>
  </el-card>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { api } from '@/utils/request'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const canManage = computed(() => userStore.hasPermission('role:list'))

const roles = ref<any[]>([])
const permissions = ref<{ id: number; code: string }[]>([])
const showDialog = ref(false)
const showPermDialog = ref(false)
const isEdit = ref(false)
const roleForm = reactive({ id: 0, name: '', description: '', isDefault: false })
const selectedPermIds = ref<(string | number)[]>([])
const currentRole = ref<any>(null)
const dialogTitle = computed(() => isEdit.value ? '编辑角色' : '新增角色')

const permTreeRef = ref<any>(null)

// 构建多层级树形数据
const treeData = computed(() => {
  const roots: any[] = []
  const nodeMap: Record<string, any> = {}
  permissions.value.forEach(p => {
    const segs = p.code.split(':')
    let path = ''
    let parentArr = roots
    segs.forEach((seg, idx) => {
      path = idx === 0 ? seg : `${path}:${seg}`
      if (!nodeMap[path]) {
        // parent nodes use string path as id; leaf nodes use numeric permission id
        nodeMap[path] = { label: seg, children: [], id: idx === segs.length - 1 ? p.id : path }
        parentArr.push(nodeMap[path])
      }
      parentArr = nodeMap[path].children
    })
  })
  return roots
})

function onTreeCheckChange() {
  if (permTreeRef.value) {
    // only include numeric permission IDs (leaf nodes)
    selectedPermIds.value = permTreeRef.value.getCheckedKeys().filter((key: string | number) => typeof key === 'number')
  }
}

async function fetchRoles() {
  try {
    const res = await api.get('/role/list')
    roles.value = res.list
  } catch {
    ElMessage.error('获取角色列表失败')
  }
}
onMounted(fetchRoles)

function openDialog(mode: 'create' | 'edit', row?: any) {
  if (!canManage.value) return
  isEdit.value = (mode === 'edit')
  if (isEdit.value && row) {
    roleForm.id = row.id
    roleForm.name = row.name
    roleForm.description = row.description
    roleForm.isDefault = row.isDefault
  } else {
    roleForm.id = 0
    roleForm.name = ''
    roleForm.description = ''
    roleForm.isDefault = false
  }
  showDialog.value = true
}

async function submitForm() {
  try {
    if (isEdit.value) {
      await api.put(`/role/${roleForm.id}`, roleForm)
    } else {
      await api.post('/role', roleForm)
    }
    ElMessage.success(isEdit.value ? '更新成功' : '创建成功')
    showDialog.value = false
    fetchRoles()
  } catch {
    ElMessage.error(isEdit.value ? '更新失败' : '创建失败')
  }
}

async function deleteRole(row: any) {
  if (!canManage.value) return
  try {
    await ElMessageBox.confirm('确定要删除该角色？', '提示', { type: 'warning' })
    await api.delete(`/role/${row.id}`)
    ElMessage.success('删除成功')
    fetchRoles()
  } catch {}
}

// 打开权限分配对话框
async function openPermDialog(row: any) {
  if (!canManage.value) return
  currentRole.value = row
  try {
    // fetch flat permissions
    const res = await api.get<{ list: any[] }>('/permission/list', { offset: 0, limit: 100 })
    permissions.value = res.list.map(p => ({ id: p.id, code: p.code }))
    // fetch already assigned
    const resp = await api.get< any[] >('/relation/role-permissions', { roleId: row.id })
    selectedPermIds.value = resp.map(rp => rp.permissionId)
  } catch {
    ElMessage.error('获取权限列表失败')
    return
  }
  showPermDialog.value = true
}

// 提交权限分配
async function submitPerms() {
  try {
    // 获取当前权限关联
    const existing = await api.get<any[]>('/relation/role-permissions', { roleId: currentRole.value.id })
    const existingIds = existing.map(rp => rp.permissionId)
    const toAdd = selectedPermIds.value.filter(id => !existingIds.includes(id))
    const toRemove = existingIds.filter(id => !selectedPermIds.value.includes(id))
    // 执行新增与删除
    await Promise.all([
      ...toAdd.map(id => api.post('/relation/role-permission', { roleId: currentRole.value.id, permissionId: id })),
      ...toRemove.map(id => api.delete('/relation/role-permission', { roleId: currentRole.value.id, permissionId: id }))
    ])
    ElMessage.success('权限分配成功')
    showPermDialog.value = false
  } catch {
    ElMessage.error('权限分配失败')
  }
}
</script>
