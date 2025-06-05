import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    redirect: '/home'
  },
  {
    path: '/home',
    name: 'Home',
    component: () => import('@/views/Home.vue'),
    meta: { 
      title: '首页',
      requiresAuth: false 
    }
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('@/views/About.vue'),
    meta: { 
      title: '关于我们',
      requiresAuth: false 
    }
  },
  {
    path: '/features',
    name: 'Features',
    component: () => import('@/views/Features.vue'),
    meta: { 
      title: '产品特性',
      requiresAuth: false 
    }
  },
  {
    path: '/pricing',
    name: 'Pricing',
    component: () => import('@/views/Pricing.vue'),
    meta: { 
      title: '价格方案',
      requiresAuth: false 
    }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/auth/Login.vue'),
    meta: { 
      title: '登录',
      requiresAuth: false 
    }
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/auth/Register.vue'),
    meta: { 
      title: '租户注册',
      requiresAuth: false 
    }
  },
  {
    path: '/dashboard',
    component: () => import('@/layouts/DashboardLayout.vue'),
    meta: { 
      requiresAuth: true 
    },
    children: [
      {
        path: '',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/Overview.vue'),
        meta: { 
          title: '工作台',
          requiresAuth: true 
        }
      },
      {
        path: 'devices',
        name: 'Devices',
        component: () => import('@/views/dashboard/Devices.vue'),
        meta: { 
          title: '设备管理',
          requiresAuth: true 
        }
      },
      {
        path: 'missions',
        name: 'Missions',
        component: () => import('@/views/dashboard/Missions.vue'),
        meta: { 
          title: '任务管理',
          requiresAuth: true 
        }
      },
      {
        path: 'users',
        name: 'Users',
        component: () => import('@/views/dashboard/Users.vue'),
        meta: { 
          title: '用户管理',
          requiresAuth: true,
          permission: 'user_manage'
        }
      },
      {
        path: 'tenant',
        name: 'TenantSettings',
        component: () => import('@/views/dashboard/TenantSettings.vue'),
        meta: { 
          title: '租户设置',
          requiresAuth: true,
          permission: 'tenant_manage'
        }
      },
      {
        path: 'profile',
        name: 'Profile',
        component: () => import('@/views/dashboard/Profile.vue'),
        meta: { 
          title: '个人中心',
          requiresAuth: true 
        }
      }
    ]
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/NotFound.vue'),
    meta: { 
      title: '页面不存在',
      requiresAuth: false 
    }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

// 路由守卫
router.beforeEach((to, from, next) => {
  // 设置页面标题
  if (to.meta?.title) {
    document.title = `${to.meta.title} - 天枢无人机智能管控平台`
  }
  
  // 检查是否需要登录
  if (to.meta?.requiresAuth) {
    const token = localStorage.getItem('token')
    if (!token) {
      next('/login')
      return
    }
  }
  
  next()
})

export default router
