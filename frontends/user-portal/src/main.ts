import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import './style.css'
import { createPinia } from 'pinia'
import { useUserStore } from '@/stores/user'

const app = createApp(App)
// 注册 Pinia
const pinia = createPinia()
app.use(pinia)
app.use(router)
app.use(ElementPlus)

// 持久化用户与权限：页面刷新时恢复登录状态
const userStore = useUserStore()
const storedUser = localStorage.getItem('userInfo')
if (storedUser) {
  try {
    const user = JSON.parse(storedUser)
    userStore.setUser(user)
    userStore.fetchPermissions()
  } catch {
    // ignore parsing errors
  }
}

app.mount('#app')
