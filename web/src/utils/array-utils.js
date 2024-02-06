/**
 * 检查是否有重复的 key
 * @param arr 数组
 * @param fieldName
 */
export function hasDuplicateKey(arr, fieldName) {
  const fieldSet = new Set()
  for (const obj of arr) {
    const fieldValue = obj[fieldName]
    if (fieldSet.has(fieldValue)) {
      return true // 有重复的 key
    }
    fieldSet.add(obj.key)
  }

  return false // 没有重复的 key
}
