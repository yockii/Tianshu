<template>
    <main>
        <div class="login-container">
            <div class="logo-row">
                <img src="@/assets/dji-logo-vector.svg" alt="DJI Logo" class="logo" />
                <t-icon name="swap" class="exchange-icon" />
                <img src="@/assets/logo-vector.svg" alt="Xjj Logo" class="logo" />
            </div>

            <div class="text-row">
                <t-typography-text strong style="font-size: 28px; color: #0582EE">天枢无人机系统</t-typography-text>
            </div>

            <div class="login-form-row">
                <t-input v-model="user.username" placeholder="请输入用户名" />
                <t-input v-model="user.password" type="password" placeholder="请输入密码" />
                <t-button theme="primary" @click="onLogin">登录</t-button>
            </div>
        </div>
    </main>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { MessagePlugin } from 'tdesign-vue-next'
import { getCloudApiInfo } from '@/api/method/sys'
import pilotBridge from '@/dji/pilot-bridge'


const isVerified = ref<boolean>(false)
const user = ref({
    username: '',
    password: ''
})

const onLogin = () => {
    if (!isVerified.value) {
        MessagePlugin.error('大疆许可证校验失败，请联系管理员处理')
        return
    }

    if (user.value.username && user.value.password) {
        // 模拟登录逻辑
        console.log('登录信息:', user.value)
        // 在这里可以添加实际的登录请求逻辑
    } else {
        console.error('用户名或密码不能为空')
    }
}

onMounted(async () => {
    await verifyLicense()
    if (!isVerified.value) {
        return
    }
})

const verifyLicense = async () => {
    const cloudApiInfo = await getCloudApiInfo()
    console.log('Cloud API Info:', cloudApiInfo)
    isVerified.value = pilotBridge.platformVerifyLicense(cloudApiInfo.appId, cloudApiInfo.appKey, cloudApiInfo.appLicense)
    if (isVerified.value) {
        MessagePlugin.success('许可证校验成功!')
    } else {
        MessagePlugin.error('大疆许可证校验失败，请联系管理员处理')
    }
}
</script>

<style scoped>
.login-container {
    width: 100vw;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    background-color: #f0f2f5;
}

.logo-row {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 20px;
    margin-top: 20px;
}

.logo {
    width: auto;
    height: 80px;
}

.exchange-icon {
    font-size: 48px;
}

.text-row {
    margin: 40px 0;
}

.login-form-row {
    /* width: 300px; */
    display: flex;
    gap: 20px;
}
</style>