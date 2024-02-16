import request from '@/utils/request'

/**
 * 备份主机创建
 * @param data
 * @returns {*}
 */
export function backupHostCreatePost(data) {
  return request({
    url: 'backup/host',
    method: 'POST',
    data
  })
}
