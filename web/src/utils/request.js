import axios from 'axios'
import { Message } from 'element-ui'
import store from '@/store'
import { getToken } from '@/utils/auth'
import { showFullScreenLoading, tryHideFullScreenLoading } from '@/utils/requestLoading'

// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 12000000 // request timeout
})
// request interceptor
service.interceptors.request.use(
  config => {
    // do something before request is sent

    // 判断是否需要加载动画
    if (config.turnOffLoading !== true) {
      showFullScreenLoading()
    }

    if (store.getters.token) {
      config.headers['Authorization'] = getToken()
    }
    return config
  },
  error => {
    // 关闭loading
    tryHideFullScreenLoading()
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
   */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    // 关闭loading
    tryHideFullScreenLoading()

    // 下载文件
    const respContentType = response.headers['content-type']
    if (respContentType === 'application/x-zip-compressed' || respContentType === 'application/octet-stream') {
      return response
    }

    const res = response.data

    // api 接口
    if (res.code !== 0 && res.status !== 'success') {
      Message({
        message: res.result || res.message || '错误',
        type: 'error',
        duration: 5 * 1000
      })
      if (res.code === '1003') {
        Message({
          message: '权限不足',
          type: 'warning',
          duration: 5 * 1000
        })
      }
      if (res.code === 50008 || res.code === 50012 || res.code === 50014) {
        // 触发这个说明认证未通过
        store.dispatch('user/resetToken').then(() => {
          location.reload()
        })
      }
      return Promise.reject(res)
    } else {
      return res
    }
  },
  error => {
    // 关闭loading
    tryHideFullScreenLoading()

    console.log('err' + error) // for debug
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

export default service
