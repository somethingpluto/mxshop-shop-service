package main

import (
	"Shop_service/user_service/global"
	"Shop_service/user_service/handler"
	"Shop_service/user_service/initialize"
	"Shop_service/user_service/proto"
	"Shop_service/user_service/util"
	"flag"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	IP := flag.String("ip", "127.0.0.1", "ip地址")
	Port := flag.Int("port", 8000, "端口号")
	flag.Parse()
	initialize.InitFileAbsPath()
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserService{})

	if global.ServiceConfig.RuntimeInfo.Mode != "debug" {
		createPort, err := util.GetFreePort()
		if err != nil {
			zap.S().Errorw("获取端口失败", "err", err.Error())
			return
		}
		*Port = createPort
		fmt.Println(*Port)
	}

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	zap.S().Infow("服务开始运行", "IP", *IP, "Port", *Port)
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	initialize.InitRegisterService()
	err = server.Serve(listen)
	if err != nil {
		panic("failed to start grpc: " + err.Error())
	}
}
