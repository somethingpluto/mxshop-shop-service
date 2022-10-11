package global

import (
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"
	"user_service/config"
)

var (
	DB            *gorm.DB
	ServiceConfig *config.ServiceConfig
	FilePath      *config.FilePathConfig
	NacosConfig   *config.NacosConfig
	Port          int
	Client        *api.Client
	ServiceID     string
)
