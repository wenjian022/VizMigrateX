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
 * 资产导出
 * @param dataSourcesType 资产类型
 * @returns {*}
 */
export function dataSourcesDeriveGet(dataSourcesType) {
  return request({
    url: `/dataSources/${dataSourcesType}/derive`,
    responseType: 'blob'
  })
}

/**
 *  资产导入模板
 * @param dataSourcesType 资产类型
 * @returns {*}
 */
export function dataSourcesImportTemplateGet(dataSourcesType) {
  return request({
    url: `/dataSources/${dataSourcesType}/import/template`,
    responseType: 'blob'
  })
}

/**
 * 资产导入
 * @param dataSourcesType 资产类型
 * @param data 文件
 * @returns {*}
 */
export function dataSourcesImportPost(dataSourcesType, data) {
  return request({
    url: `/dataSources/${dataSourcesType}/import`,
    method: 'POST',
    data
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

// 数据源-选择器数据格式返回
export function dataSourcesDatabasesSelectGet() {
  return request({
    url: 'dataSource/databases/select'
  })
}

/**
 * 获取数据库资产-库信息
 * @param dataSourceID
 * @returns {*}
 */
export function dataSourceDatabasesInformationGet(dataSourceID) {
  return request({
    url: `dataSource/database/information/${dataSourceID}`,
    method: 'GET'
  })
}
