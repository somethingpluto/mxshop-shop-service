package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"goods_service/config"
)

func main() {
	configFileName := "D:\\C_Back\\Go\\Shop_service\\goods_service\\config-debug.yaml"
	fmt.Println(configFileName)
	v := viper.New()
	v.SetConfigFile(configFileName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	nacosConfig := config.NacosConfig{}
	err = v.Unmarshal(&nacosConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(nacosConfig)
	// 连接 nacos
	serviceConfig := []constant.ServerConfig{
		{
			IpAddr: nacosConfig.Host,
			Port:   uint64(nacosConfig.Port),
		},
	}
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serviceConfig,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacosConfig.Dataid,
		Group:  nacosConfig.Group,
	})
	if err != nil {
		panic(err)
	}
	serverConfig := config.ServiceConfig{}
	err = json.Unmarshal([]byte(content), &serverConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)
	fmt.Printf("%#v", serverConfig)
}
