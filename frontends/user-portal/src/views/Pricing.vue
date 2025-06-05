<template>
  <el-main>
    <el-card>
      <h2>价格方案</h2>
      <el-row :gutter="24">
        <el-col :span="8">
          <el-card shadow="hover">
            <h3>基础版</h3>
            <p>适合小型团队/试用</p>
            <p>免费</p>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="hover">
            <h3>专业版</h3>
            <p>适合成长型企业</p>
            <p>￥199/月</p>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="hover">
            <h3>旗舰版</h3>
            <p>适合大中型企业/定制</p>
            <p>￥599/月</p>
          </el-card>
        </el-col>
      </el-row>
    </el-card>
  </el-main>
</template>

<script setup lang="ts">
</script>
