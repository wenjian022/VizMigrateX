package tool

// PageQueryStruct 分页
type PageQueryStruct struct {
	Size int `form:"size,default=10"` // 每页显示的数量
	Page int `form:"page,default=1"`  // 当前页码
}

// PagingDataResStruct  可分页数据的返回
type PagingDataResStruct struct {
	Data  interface{} `json:"data"`           // 分页的数据内容
	Total int         `json:"total"`          // 全部的页码数量
	Size  int         `json:"size,omitempty"` // 每页显示的数量
	Page  int         `json:"page,omitempty"` // 当前页码
}
