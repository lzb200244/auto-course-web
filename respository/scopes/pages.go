package scopes

import "gorm.io/gorm"

/*
Created by 斑斑砖 on 2023/9/8.
	Description：分页封装 Scope
*/

const (
	Size = 20
)

// Paginate 分页封装
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > Size:
			pageSize = Size
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
