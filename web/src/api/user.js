import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/auth/login',
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
    url: '/auth/logout',
    method: 'post'
  })
}

/**
 * 添加用户
 * @param data
 * @returns {*}
 * @constructor
 */
export function usersAddPost(data) {
  return request({
    url: '/user/add',
    method: 'POST',
    data
  })
}

/**
 * 查看用户列表
 * @param params
 * @returns {*}
 */
export function usersGet(params) {
  return request({
    url: '/users',
    params: { params }
  })
}

/**
 * 删除用户
 * @param id
 * @returns {*}
 */
export function usersDle(id) {
  return request({
    url: `/user/del/${id}`,
    method: 'DELETE'
  })
}
