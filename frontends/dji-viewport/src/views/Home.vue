<template>
    <main>
        <p>{{connectInfo}}</p>
       <p> {{text}}</p>
        <p>controller sn: {{ controllerSn }}</p>
    </main>
</template>

<script lang="ts" setup>
import { getConnectInfo } from '@/api/method/sys';
import pilotBridge from '@/dji/pilot-bridge';
import { useUserStore } from '@/stores/user';
import { type ConnectInfo } from '@/types/sys-info';
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router';

const text = ref('Loading...')
const connectInfo = ref<ConnectInfo|null>(null)
const controllerSn = ref<string|null>(null)

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
    
    connectInfo.value = result

    controllerSn.value = pilotBridge.getRemoteControllerSN()

    window.thingConnectCallback = thingConnectCallback

    pilotBridge.loadModule('thing', {
        host: result.mqttTcpAddr,
        username: result.mqttUsername,
        password: result.mqttPassword,
        connectCallback: 'thingConnectCallback',
    })
})

const thingConnectCallback = async (success: boolean) => {
    console.log('thingConnectCallback', success)
    if (success) {
        // 连接成功
        console.log('Thing Connect Success')
        text.value = 'Thing Connect Success'
    } else {
        // 连接失败
        console.error('Thing Connect Failed', success)
        text.value = 'Thing Connect Failed: ' + success
    }
}
</script>