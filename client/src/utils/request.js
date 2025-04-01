import axios from 'axios'
import { Message } from '@arco-design/web-vue'
import { getToken } from '@/utils/auth'

const service = axios.create({
  baseURL: import.meta.env.VITE_APP_BASE_API,
  timeout: 10000
})

// 请求拦截器
service.interceptors.request.use(
  config => {
    const token = getToken()
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// 响应拦截器
service.interceptors.response.use(
  response => {
    const res = response.data
    
    if (res.code !== 0) {
      Message.error({
        content: res.message || '请求错误',
        duration: 5 * 1000
      })
      return Promise.reject(new Error(res.message || '请求错误'))
    }
    return res.result
  },
  error => {
    Message.error({
      content: error.message || '请求错误',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

export default service
