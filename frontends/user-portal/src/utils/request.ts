import axios from 'axios'
import { ElMessage } from 'element-plus'

// 创建axios实例
const request = axios.create({
  baseURL: '/api',
  timeout: 10000,
})

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    // 添加token
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    
    // 添加租户ID
    const tenantId = localStorage.getItem('tenantId')
    if (tenantId) {
      config.headers['X-Tenant-ID'] = tenantId
    }
    
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  (response) => {
    const { data } = response

    // special case: blob
    if (response.config.responseType === 'blob') {
      return response
    }

    // wrapped response
    if (data && typeof data.code === 'number') {
      if (data.code === 200) {
        return data.data
      } else if (data.code === 401) {
        // unauthorized
        localStorage.removeItem('token')
        localStorage.removeItem('userInfo')
        window.location.href = '/login'
        return Promise.reject(new Error('未授权'))
      } else {
        ElMessage.error(data.message || '请求失败')
        return Promise.reject(new Error(data.message || '请求失败'))
      }
    }
    // raw response
    return data
  },
  (error) => {
    console.error('请求错误:', error)

    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      localStorage.removeItem('userInfo')
      window.location.href = '/login'
    } else {
      ElMessage.error(error.message || '网络错误')
    }

    return Promise.reject(error)
  }
)

export default request

// 通用API方法
export const api = {
  get: <T = any>(url: string, params?: any): Promise<T> => request.get(url, { params }),
  post: <T = any>(url: string, data?: any): Promise<T> => request.post(url, data),
  put: <T = any>(url: string, data?: any): Promise<T> => request.put(url, data),
  // delete can send request body via data property
  delete: <T = any>(url: string, data?: any): Promise<T> => request.delete(url, { data })
}
