import request from '@/utils/request'

/**
 * 数据源连接测试
 * @param data
 * @returns {*}
 */
export function dataSourcesConnectionTestPost(data) {
  return request({
    url: 'dataSource/testConnection',
    method: 'POST',
    data
  })
}

/**
 * 数据源连接测试
 * @param dataSourcesID
 * @returns {*}
 */
export function dataSourcesConnectionTestIDPost(dataSourcesID) {
  return request({
    url: `dataSource/testConnection/${dataSourcesID}`,
    method: 'POST'
  })
}

/**
 * 创建数据源
 * @param data
 * @returns {*}
 */
export function dataSourcesDatabasesCratePost(data) {
  return request({
    url: '/dataSource',
    method: 'POST',
    data
  })
}

/**
 * 获取数据源详情
 * @param dataSourcesId
 */
export function dataSourcesInfoGet(dataSourcesId) {
  return request({
    url: `dataSource/${dataSourcesId}`,
    method: 'GET'
  })
}

/**
 * 修改数据源
 * @param data
 * @param dataSourcesId
 * @returns {*}
 */
export function dataSourcesDatabasesPut(data, dataSourcesId) {
  return request({
    turnOffLoading: true,
    url: `dataSource/${dataSourcesId}`,
    method: 'PUT',
    data
  })
}

/**
 * 删除数据源
 * @param dataSourcesId 资产id
 * @returns {*}
 */
export function dataSourcesDelete(dataSourcesId) {
  return request({
    url: `dataSource/${dataSourcesId}`,
    method: 'DELETE',
    data: { 'dataSourcesId': dataSourcesId }
  })
}

/**
 * 获取所有数据源
 * @returns {*}
 */
export function dataSourcesDatabasesGet(params) {
  return request({
    url: '/dataSource',
    params: params
  })
}

/**
 *
 * @param params
 * @returns {*}
 */
export function dataSourcesEnvListGet(params) {
  return request({
    url: '/dataSource/env',
    params: params
  })
}

/**
 *
 * @param params
 * @returns {*}
 */
export function dataSourcesLabelListGet(params) {
  return request({
    url: '/dataSource/label',
    params: params
  })
}

/**
 *
 * @param data
 * @returns {*}
 */
export function dataSourcesLabelPost(data) {
  return request({
    url: '/dataSource/label',
    method: 'POST',
    data
  })
}

/**
 *
 * @param data
 * @returns {*}
 */
export function dataSourcesEnvPost(data) {
  return request({
    url: '/dataSource/env',
    method: 'POST',
    data
  })
}
