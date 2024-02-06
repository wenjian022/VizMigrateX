import { Loading } from 'element-ui'

let loading
let needLoadingRequestCount = 0

function startLoading() {
  loading = Loading.service({
    lock: true,
    text: '数据加载中，请稍后……',
    background: 'rgba(0, 0, 0, 0.5)'
  })
}

function endLoading() {
  // 延迟500ms，防止网速特快加载中画面一闪而过
  setTimeout(function() {
    if (loading) loading.close()
  }, 500)
}

// 打开loading
export function showFullScreenLoading() {
  if (needLoadingRequestCount === 0) {
    startLoading()
  }
  needLoadingRequestCount++
}

// 关闭loading
export function tryHideFullScreenLoading() {
  if (needLoadingRequestCount <= 0) return
  needLoadingRequestCount--
  if (needLoadingRequestCount === 0) {
    endLoading()
  }
}
