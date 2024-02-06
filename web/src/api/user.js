import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/user/auth/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/user/info',
    method: 'get'
    // params: { token }
  })
}

export function logout() {
  return request({
    url: '/user/auth/logout',
    method: 'post'
  })
}

/**
 * 查看是否需要用户初始化
 * @returns {*}
 */
export function userInitializationGet() {
  return request({
    url: '/user/initialization'
  })
}

/**
 * 初始化用户
 * @param data
 * @returns {*}
 */
export function userInitializationPost(data) {
  return request({
    url: '/user/initialization',
    method: 'post',
    data
  })
}
