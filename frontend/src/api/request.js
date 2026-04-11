import axios from 'axios'
import { ElMessage, ElLoading } from 'element-plus'

// 创建 axios 实例
const request = axios.create({
  // API 请求基地址
  // 开发环境使用 localhost:8000，生产环境可配置为实际服务器地址
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8000',
  // 请求超时时间
  timeout: 15000,
  // 请求头
  headers: {
    'Content-Type': 'application/json',
  },
})
let loadingInstance = null;
let loadingCount = 0;

const showLoading = () => {
  if (loadingCount === 0) {
    loadingInstance = ElLoading.service({
      lock: true, // 锁定屏幕
      text: '加载中...', // 加载文字
      spinner: 'el-icon-loading', // 加载图标
    })
  }
  loadingCount++
}
const hideLoading = () => {
  loadingCount--
  if (loadingCount <= 0 && loadingInstance) {
    loadingInstance?.close()
    loadingInstance = null
  }
}

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    showLoading();
    // 添加 token 等认证信息
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    // 添加公共参数
    config.headers['X-Request-ID'] = generateUUID()
    return config
  },
  (error) => {
    hideLoading();
    ElMessage.error('请求配置错误')
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  (response) => {
    hideLoading();
    const res = response.data
    const { code, message, data } = res

    // 根据业务状态码处理
    if (code === 0 || code === 200) {
      return data
    }

    // 其他业务错误
    ElMessage.error(message || '请求失败')
    return Promise.reject(new Error(message || 'Request failed'))
  },
  (error) => {
    hideLoading();
    // HTTP 错误处理
    if (error.response) {
      const { status, data } = error.response
      switch (status) {
        case 401:
          ElMessage.error('登录已过期，请重新登录')
          // 清除 token 并跳转到登录页
          localStorage.removeItem('token')
          window.location.href = '/login'
          break
        case 403:
          ElMessage.error('没有权限进行此操作')
          break
        case 404:
          ElMessage.error('请求的资源不存在')
          break
        case 500:
          ElMessage.error('服务器内部错误，请稍后再试')
          break
        case 502:
          ElMessage.error('网关错误，请稍后再试')
          break
        case 503:
          ElMessage.error('服务暂不可用，请稍后再试')
          break
        case 504:
          ElMessage.error('网关超时，请稍后再试')
          break
        default:
          ElMessage.error(data?.message || `请求失败 (HTTP ${status})`)
      }
    } else if (error.request) {
      ElMessage.error('网络连接失败，请检查网络')
    } else {
      ElMessage.error(error.message || '请求配置错误')
    }
    return Promise.reject(error)
  }
)

// 生成 UUID
function generateUUID() {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
    const r = Math.random() * 16 | 0
    const v = c === 'x' ? r : (r & 0x3 | 0x8)
    return v.toString(16)
  })
}

export default request
