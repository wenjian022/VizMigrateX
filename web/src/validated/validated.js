/**
 * 校验IP地址是否合格
 * @param rule
 * @param value
 * @param callback
 * @constructor
 */
export function validateIP(rule, value, callback) {
  if (value === '' || typeof value === 'undefined' || value == null) {
    callback(new Error('请输入正确的IP地址'))
  } else {
    const reg = /^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/
    if ((!reg.test(value)) && value !== '') {
      callback(new Error('请输入正确的IP地址'))
    } else {
      callback()
    }
  }
}

/**
 * 端口号校验
 * @param rule
 * @param value
 * @param callback
 */
export function validatorPort(rule, value, callback) {
  if (value === '' || typeof value === 'undefined' || value == null || isNaN(value)) {
    callback(new Error('请输入正确的端口号'))
  } else {
    if (value > 1 && value < 65535) {
      callback()
    } else {
      callback(new Error('请输入正确的端口号'))
      callback()
    }
  }
}
