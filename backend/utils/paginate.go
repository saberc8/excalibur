package utils

import "gorm.io/gorm"

// Paginate 定义分页结构体，绑定到global.AquilaDb
func Paginate(pageNum int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageSize == 0 {
			pageSize = 10
		}
		if pageNum == 0 {
			pageNum = 1
		}
		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
