import request from '@/utils/request'

/**
 * 数据库连接测试
 * @param data
 * @returns {*}
 * @constructor
 */
export function assetsConnectionTestPost(data) {
  return request({
    url: 'asset/connection/test',
    method: 'POST',
    data
  })
}

/**
 * 创建MySQL资产
 * @param data
 * @returns {*}
 */
export function assetsDatabasesCratePost(data) {
  return request({
    url: 'asset/databases',
    method: 'POST',
    data
  })
}

/**
 * 修改MySQL资产
 * @param data
 * @returns {*}
 */
export function assetsDatabasesPut(data, id) {
  return request({
    url: `asset/database/${id}`,
    method: 'PUT',
    data
  })
}

/**
 * 删除资产
 * @param assetsId 资产id
 * @returns {*}
 */
export function assetsDelete(assetsId) {
  return request({
    url: 'assets',
    method: 'DELETE',
    data: { 'assetsId': assetsId }
  })
}

/**
 * 资产导出
 * @param assetsType 资产类型
 * @returns {*}
 */
export function assetsDeriveGet(assetsType) {
  return request({
    url: `/assets/${assetsType}/derive`,
    responseType: 'blob'
  })
}

/**
 *  资产导入模板
 * @param assetsType 资产类型
 * @returns {*}
 */
export function assetsImportTemplateGet(assetsType) {
  return request({
    url: `/assets/${assetsType}/import/template`,
    responseType: 'blob'
  })
}

/**
 * 资产导入
 * @param assetsType 资产类型
 * @param data 文件
 * @returns {*}
 */
export function assetsImportPost(assetsType, data) {
  return request({
    url: `/assets/${assetsType}/import`,
    method: 'POST',
    data
  })
}

/**
 * 获取所有数据库资产
 * @returns {*}
 */
export function assetsDatabasesGet(params) {
  return request({
    url: '/asset/database/list',
    params: params
  })
}

// 资产数据库-选择器数据格式返回
export function assetsDatabasesSelectGet() {
  return request({
    url: 'asset/databases/select'
  })
}

/**
 * 获取数据库资产-库信息
 * @param assetID
 * @returns {*}
 */
export function assetDatabasesInformationGet(assetID) {
  return request({
    url: `asset/database/information/${assetID}`,
    method: 'GET'
  })
}
