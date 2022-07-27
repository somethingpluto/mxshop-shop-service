package global

import (
	"Shop_service/user_service/config"
	"gorm.io/gorm"
)

var DB *gorm.DB
var ServiceConfig = &config.ServiceConfig{}
