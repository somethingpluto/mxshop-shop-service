package global

import (
	"Shop_service/user_service/config"
	"gorm.io/gorm"
)

//  数据库连接
var DB *gorm.DB

//  配置文件
var ServiceConfig = &config.ServiceConfig{}

var FilePath = &config.FilePathConfig{}
