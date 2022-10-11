package main

import (
	"flag"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"user_service/initialize"
	"user_service/mode"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 8000, "端口号")
	Mode := flag.String("mode", "release", "mode debug 本地调试/ release 服务主从")
	flag.Parse()
	initialize.InitFileAbsPath()
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	server := grpc.NewServer()
	if *Mode == "debug" {
		zap.S().Warnf("debug本地调试模式 \n")
		mode.DebugMode(server, *IP, *Port)
	} else if *Mode == "release" {
		zap.S().Warnf("release服务注册模式 \n")
		mode.ReleaseMode(server, *IP)
	}
}
