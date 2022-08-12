package global

import (
	"goods_service/config"
	"gorm.io/gorm"
)

var (
	DB            *gorm.DB
	FilePath      *config.FilePathConfig
	ServiceConfig *config.ServiceConfig
	NacosConfig   *config.NacosConfig
)
