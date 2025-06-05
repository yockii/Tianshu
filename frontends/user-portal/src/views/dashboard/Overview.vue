<template>
  <el-card>
    <h2>工作台</h2>
    <p>欢迎使用天枢无人机智能管控平台！</p>
    <el-row :gutter="24">
      <el-col :span="8">
        <el-statistic title="设备总数" :value="128" />
      </el-col>
      <el-col :span="8">
        <el-statistic title="在线设备" :value="96" />
      </el-col>
      <el-col :span="8">
        <el-statistic title="今日任务" :value="32" />
      </el-col>
    </el-row>
  </el-card>
</template>

<script setup lang="ts">
</script>
