package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"goods_service/global"
)

func InitConfig() {
	configFileName := fmt.Sprintf(global.FilePath.ConfigFile)
	v := viper.New()
	v.SetConfigFile(configFileName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
		return
	}
	err = v.Unmarshal(&global.ServiceConfig)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("%#v \n", global.ServiceConfig)
	fmt.Println("配置文件加载成功")
}
