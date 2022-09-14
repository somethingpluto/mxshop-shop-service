package global

import (
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"
	"user_service/config"
)

var (
	DB                *gorm.DB
	UserServiceConfig *config.UserServiceConfig
	FilePath          *config.FilePathConfig
	NacosConfig       *config.NacosConfig
	Port              int
	Client            *api.Client
	ServiceID         string
)
