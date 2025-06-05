// 用户类型
export interface User {
  id: string
  tenantId: string
  username: string
  email: string
  phone?: string
  name: string
  avatar?: string
  status: number
  roles: Role[]
  permissions: string[]
  createdAt: string
  updatedAt: string
}

// 租户类型
export interface Tenant {
  id: string
  name: string
  code: string
  domain?: string
  logo?: string
  primaryColor?: string
  secondaryColor?: string
  welcomeMessage?: string
  status: number
  expiredAt?: string
  createdAt: string
  updatedAt: string
}

// 角色类型
export interface Role {
  id: string
  tenantId: string
  name: string
  code: string
  description?: string
  permissions: Permission[]
  createdAt: string
  updatedAt: string
}

// 权限类型
export interface Permission {
  id: string
  name: string
  code: string
  type: string
  resource?: string
  action?: string
  description?: string
}

// 设备类型
export interface Device {
  id: string
  tenantId: string
  name: string
  type: string
  model?: string
  serialNumber: string
  status: 'online' | 'offline' | 'error' | 'maintenance'
  location?: {
    latitude: number
    longitude: number
    altitude?: number
  }
  battery?: number
  temperature?: number
  lastHeartbeat?: string
  createdAt: string
  updatedAt: string
}

// 任务类型
export interface Mission {
  id: string
  tenantId: string
  name: string
  description?: string
  deviceId: string
  routeId?: string
  status: 'pending' | 'running' | 'completed' | 'failed' | 'cancelled'
  startTime?: string
  endTime?: string
  progress: number
  createdAt: string
  updatedAt: string
}

// API响应类型
export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
  timestamp: string
}

// 分页类型
export interface Pagination {
  page: number
  size: number
  total: number
}

export interface PaginatedResponse<T = any> extends ApiResponse<T[]> {
  pagination: Pagination
}

// 表单验证规则类型
export interface FormRule {
  required?: boolean
  message?: string
  min?: number
  max?: number
  pattern?: RegExp
  validator?: (rule: any, value: any, callback: any) => void
}

// 菜单项类型
export interface MenuItem {
  id: string
  name: string
  path: string
  icon?: string
  component?: string
  children?: MenuItem[]
  meta?: {
    title: string
    requiresAuth?: boolean
    roles?: string[]
    permissions?: string[]
  }
}

// 主题配置类型
export interface ThemeConfig {
  primaryColor: string
  secondaryColor: string
  logo?: string
  title: string
  favicon?: string
}
