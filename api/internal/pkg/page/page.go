package page

import "gorm.io/gorm"

// PagingQueryStruct 分页
type PagingQueryStruct struct {
	Size int `form:"size,default=10" json:"size"` // 每页显示的数量
	Page int `form:"page,default=1" json:"page"`  // 当前页码
}

// PagingDataResStruct  可分页数据的返回
type PagingDataResStruct struct {
	Data  interface{} `json:"data"`  // 分页的数据内容
	Total int64       `json:"total"` // 全部的页码数量

	PagingQueryStruct
}

func Paginate(req *PagingQueryStruct) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := req.Page
		if page == 0 {
			page = 1
		}
		pageSize := req.Size
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
