package util

type Code struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Result  interface{} `json:"result"`
}

// UtilSuccess 请求成功的输出
func UtilSuccess(result interface{}) Code {
	return Code{
		Code:    0,
		Message: "ok",
		Result:  result,
	}
}

// UtilError 请求错误的输出
func UtilError(result interface{}) Code {
	return Code{
		Code:    1,
		Message: "请求错误",
		Result:  result,
	}
}

// QueryUriId
//
//	@Description: uriId 校验
type QueryUriId struct {
	Id int64 `uri:"id" form:"id" binding:"required"`
}
