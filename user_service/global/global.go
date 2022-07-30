package global

import (
	"Shop_service/user_service/config"
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"
)

var (
	DB            *gorm.DB
	ServiceConfig = &config.ServiceConfig{}
	FilePath      = &config.FilePathConfig{}
	Port          *int
	Client        *api.Client
	ServiceID     string
)
