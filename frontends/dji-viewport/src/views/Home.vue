<template>
    <main>{{text}}</main>
</template>

<script lang="ts" setup>
import { getConnectInfo } from '@/api/method/sys';
import pilotBridge from '@/dji/pilot-bridge';
import { useUserStore } from '@/stores/user';
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router';

const text = ref('Loading...')

const router = useRouter()
const userStore = useUserStore()

onMounted(async () => {
    pilotBridge.onBackClick(() => {
        router.replace({
            name: 'Login'
        })
    })
    pilotBridge.onStopPlatform(() => {
        userStore.logout()
    })

    const result = await getConnectInfo()
    

    window.thingConnectCallback = thingConnectCallback

    pilotBridge.loadModule('thing', {
        host: result.mqttWsAddr,
        username: result.mqttUsername,
        password: result.mqttPassword,
        connectCallback: 'thingConnectCallback',
    })
})

const thingConnectCallback = async (arg: any) => {
    console.log('thingConnectCallback', arg)
    if (arg.code === 0) {
        // 连接成功
        console.log('Thing Connect Success')
        text.value = 'Thing Connect Success'
    } else {
        // 连接失败
        console.error('Thing Connect Failed', arg.message)
        text.value = 'Thing Connect Failed: ' + arg.message
    }
}
</script>