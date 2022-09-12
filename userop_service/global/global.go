package global

import (
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"
	"userop_service/config"
)

var (
	FilePath      *config.FilePathConfig
	NacosConfig   *config.NacosConfig
	ServiceConfig *config.ServiceConfig
	DB            *gorm.DB
	Client        *api.Client
	FreePort      int
	ServiceID     string
)
