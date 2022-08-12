package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"goods_service/config"
	"goods_service/global"
)

// InitConfig
// @Description: 初始化配置
//
func InitConfig() {
	// 获得配置文件路径
	configFileName := fmt.Sprintf(global.FilePath.ConfigFile)
	// 生成viper
	v := viper.New()
	// 指定配置文件
	v.SetConfigFile(configFileName)
	err := v.ReadInConfig()
	if err != nil {
		zap.S().Errorw("viper 读取配置文件失败", "err", err.Error())
		return
	}
	global.NacosConfig = &config.NacosConfig{}
	err = v.Unmarshal(global.NacosConfig)
	if err != nil {
		zap.S().Errorw("配置文件 解析到global.serviceConfig失败", "err", err.Error())
		return
	}
	zap.S().Infof("%#v", global.NacosConfig)
	// 连接nacos
	sConfig := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   uint64(global.NacosConfig.Port),
		},
	}

	nacosLogDir := fmt.Sprintf("%s/%s/%s", global.FilePath.LogFile, "nacos", "log")
	nacosCacheDir := fmt.Sprintf("%s/%s/%s", global.FilePath.LogFile, "nacos", "cache")
	cConfig := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              nacosLogDir,
		CacheDir:            nacosCacheDir,
		LogLevel:            "debug",
	}
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sConfig,
		"clientConfig":  cConfig,
	})
	if err != nil {
		zap.S().Errorw("客户端连接失败", "err", err.Error())
		return
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.Dataid,
		Group:  global.NacosConfig.Group,
	})
	if err != nil {
		zap.S().Errorw("client.GetConfig读取文件失败", "err", err.Error())
		return
	}
	global.ServiceConfig = &config.ServiceConfig{}
	err = json.Unmarshal([]byte(content), global.ServiceConfig)
	if err != nil {
		zap.S().Errorw("读取的配置content解析到global.serviceConfig失败", "err", err.Error())
		return
	}
	zap.S().Infof("nacos配置拉取成功 %#v", global.ServiceConfig)
}
