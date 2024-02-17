import request from '@/utils/request'

/**
 * 数据复制任务列表
 * @param params
 * @returns {*}
 */
export function dataReplicationTaskListGet(params) {
  return request({
    url: 'dataReplication/task/list',
    method: 'GET',
    params
  })
}

/**
 * 添加数据复制任务
 * @param data
 * @returns {*}
 */
export function dataReplicationTaskPost(data) {
  return request({
    url: 'dataReplication/task',
    method: 'POST',
    data
  })
}

/**
 * 获取任务详情
 * @param taskID
 * @returns {*}
 */
export function dataReplicationTaskInfoGet(taskID) {
  return request({
    url: `dataReplication/task/${taskID}`,
    method: 'GET'
  })
}

/**
 * 修改数据复制任务
 * @param step
 * @param taskID
 * @param data
 * @returns {*}
 */
export function dataReplicationTaskPut(step, taskID, data) {
  return request({
    url: `dataReplication/task/${step}/${taskID}`,
    method: 'PUT',
    data
  })
}
