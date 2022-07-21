package util

import (
	"Shop_service/user_service/global"
	"gorm.io/gorm"
)

// Paginate
// @Description: 分页
// @param page
// @param pageSize
// @return func(db *gorm.DB) *gorm.DB
//
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return global.DB.Offset(offset).Limit(pageSize)
	}
}
