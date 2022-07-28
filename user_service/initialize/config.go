package initialize

import (
	"Shop_service/user_service/global"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// InitConfig
// @Description: 初始化配置
//
func InitConfig() {
	configFileName := fmt.Sprintf(global.FilePath.ConfigFile)
	v := viper.New()
	v.SetConfigFile(configFileName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(global.ServiceConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(global.ServiceConfig)
	zap.S().Infow("config文件加载成功")
}
