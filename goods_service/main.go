package main

import (
	"flag"
	"go.uber.org/zap"
	"goods_service/global"
	"goods_service/initialize"
	"goods_service/mode"
	"google.golang.org/grpc"
)

func main() {
	IP := flag.String("ip", "127.0.0.1", "ip地址")
	Port := flag.Int("port", 8000, "port端口号")
	flag.Parse()
	global.FreePort = *Port
	// 初始化文件路径
	initialize.InitFileAbsPath()
	// 初始化配置文件
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库
	initialize.InitDB()
	server := grpc.NewServer()
	if global.ServiceConfig.Mode == "debug" {
		zap.S().Warnln("debug本地调试模式")
		mode.DebugMode(server, *IP)
	} else if global.ServiceConfig.Mode == "release" {
		zap.S().Warnln("online 服务注册模式")
		mode.OnlineMode(server, *IP)
	}
}
