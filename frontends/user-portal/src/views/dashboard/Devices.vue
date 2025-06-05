<template>
  <el-card>
    <h2>设备管理</h2>
    <el-table :data="[]">
      <el-table-column prop="id" label="设备ID" />
      <el-table-column prop="name" label="设备名称" />
      <el-table-column prop="type" label="类型" />
      <el-table-column prop="status" label="状态" />
      <el-table-column prop="createdAt" label="注册时间" />
    </el-table>
  </el-card>
</template>

<script setup lang="ts">
</script>
