import axios from 'axios'

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

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    // 可在此处添加 token 等认证信息
    // const token = localStorage.getItem('token')
    // if (token) {
    //   config.headers.Authorization = `Bearer ${token}`
    // }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  (response) => {
    const { code, message, data } = response.data
    // 根据业务状态码处理
    if (code === 0 || code === 200) {
      return data
    }
    // 其他业务错误
    console.error(`[API Error] ${message || 'Unknown error'}`)
    return Promise.reject(new Error(message || 'Request failed'))
  },
  (error) => {
    // HTTP 错误处理
    if (error.response) {
      const { status, data } = error.response
      switch (status) {
        case 401:
          console.error('[API Error] Unauthorized - please login')
          // 可触发跳转到登录页
          break
        case 403:
          console.error('[API Error] Forbidden - no permission')
          break
        case 404:
          console.error('[API Error] Not Found')
          break
        case 500:
          console.error('[API Error] Server Error')
          break
        default:
          console.error(`[API Error] HTTP ${status}: ${data?.message || error.message}`)
      }
    } else if (error.request) {
      console.error('[API Error] No response received - please check network')
    } else {
      console.error('[API Error]', error.message)
    }
    return Promise.reject(error)
  }
)

export default request
