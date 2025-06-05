<template>
  <el-card>
    <h2>任务管理</h2>
    <el-table :data="[]">
      <el-table-column prop="id" label="任务ID" />
      <el-table-column prop="name" label="任务名称" />
      <el-table-column prop="status" label="状态" />
      <el-table-column prop="deviceId" label="设备ID" />
      <el-table-column prop="startTime" label="开始时间" />
      <el-table-column prop="endTime" label="结束时间" />
    </el-table>
  </el-card>
</template>

<script setup lang="ts">
</script>
