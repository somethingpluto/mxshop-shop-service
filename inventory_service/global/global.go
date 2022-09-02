package global

import (
	"gorm.io/gorm"
	"inventory_service/config"
)

var (
	FilePath      *config.FilePathConfig
	NacosConfig   *config.NacosConfig
	ServiceConfig *config.ServiceConfig
	DB            *gorm.DB
)
