package global

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"
	"inventory_service/config"
)

var (
	FilePath      *config.FilePathConfig
	NacosConfig   *config.NacosConfig
	ServiceConfig *config.ServiceConfig
	DB            *gorm.DB
	Client        *api.Client
	FreePort      int
	ServiceID     string
	Redsync       *redsync.Redsync
)
