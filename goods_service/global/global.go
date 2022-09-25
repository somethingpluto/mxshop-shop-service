package global

import (
	"github.com/hashicorp/consul/api"
	"github.com/olivere/elastic/v7"
	"goods_service/config"
	"gorm.io/gorm"
)

var (
	DB            *gorm.DB
	FilePath      *config.FilePathConfig
	ServiceConfig *config.ServiceConfig
	NacosConfig   *config.NacosConfig
	FreePort      int
	Client        *api.Client
	ServiceID     string
	EsClient      *elastic.Client
)
