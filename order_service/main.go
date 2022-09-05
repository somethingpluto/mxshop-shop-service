package main

import (
	"flag"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"order_service/initialize"
	"order_service/mode"
)

func main() {
	IP := flag.String("ip", "127.0.0.1", "ip地址：服务启动ip地址")
	Port := flag.Int("port", 8000, "port端口号：服务启动端口号")
	Mode := flag.String("mode", "debug", "mode启动模式：debug 本地调试/release 服务注册")
	flag.Parse()
	initialize.InitFileAbsPath()
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
	initialize.InitOtherService()
	server := grpc.NewServer()
	if *Mode == "debug" {
		zap.S().Warnf("debug本地调试模式 \n")
		mode.DebugMode(server, *IP, *Port)
	} else if *Mode == "release" {
		zap.S().Warnf("release服务注册模式 \n")
		mode.ReleaseMode(server, *IP)
	}
}
