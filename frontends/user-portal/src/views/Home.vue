<template>
  <div class="home-page">
    <el-header height="80px" class="home-header">
      <img src="/vite.svg" alt="logo" class="logo" />
      <h1>天枢无人机智能管控平台</h1>
      <el-button type="primary" @click="$router.push('/login')">登录</el-button>
      <el-button @click="$router.push('/register')">注册</el-button>
    </el-header>
    <el-main>
      <section class="hero">
        <h2>多租户SaaS · 智能 · 安全 · 高效</h2>
        <p>为无人机行业提供一站式智能管控解决方案</p>
        <el-button type="primary" size="large" @click="$router.push('/features')">了解产品特性</el-button>
      </section>
      <section class="features-preview">
        <el-row :gutter="24">
          <el-col :span="8">
            <el-card shadow="hover">
              <h3>多租户隔离</h3>
              <p>支持租户自主管理，数据隔离，安全可靠</p>
            </el-card>
          </el-col>
          <el-col :span="8">
            <el-card shadow="hover">
              <h3>智能航线与任务</h3>
              <p>可视化航线规划，任务自动下发与监控</p>
            </el-card>
          </el-col>
          <el-col :span="8">
            <el-card shadow="hover">
              <h3>实时数据与视频</h3>
              <p>设备状态、视频流、告警信息一屏掌控</p>
            </el-card>
          </el-col>
        </el-row>
      </section>
    </el-main>
    <el-footer height="60px" class="home-footer">
      <span>© 2025 天枢无人机智能管控平台</span>
    </el-footer>
  </div>
</template>

<script setup lang="ts">
// 首页无需特殊逻辑
</script>

<style scoped lang="scss">
.home-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #fff;
  box-shadow: 0 2px 8px #f0f1f2;
  .logo {
    height: 48px;
    margin-right: 16px;
  }
}
.hero {
  text-align: center;
  margin: 60px 0 40px 0;
  h2 {
    font-size: 2.2rem;
    margin-bottom: 12px;
  }
  p {
    color: #888;
    margin-bottom: 24px;
  }
}
.features-preview {
  margin: 40px 0;
}
.home-footer {
  text-align: center;
  color: #888;
  background: #fafbfc;
}
</style>
