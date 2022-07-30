package main

import (
	"Shop_service/user_service/global"
	"Shop_service/user_service/handler"
	"Shop_service/user_service/initialize"
	"Shop_service/user_service/proto"
	"Shop_service/user_service/util"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	IP := flag.String("ip", "127.0.0.1", "ip地址")
	global.Port = flag.Int("port", 8000, "端口号")
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
		*global.Port = createPort
		fmt.Println(global.Port)
	}

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *global.Port))
	zap.S().Infow("服务开始运行", "IP", *IP, "Port", *global.Port)
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	initialize.InitRegisterService()
	go func() {
		err = server.Serve(listen)
		if err != nil {
			panic("failed to start grpc: " + err.Error())
		}
	}()

	// TODO:注册服务失败报错
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("serviceID", global.ServiceID)
	err = global.Client.Agent().ServiceDeregister(global.ServiceID)
	if err != nil {
		zap.S().Errorw("global.Client.Agent().ServiceDeregister 失败", "err", err.Error())
		return
	}
	zap.S().Infow("服务注销程", "serviceID", global.ServiceID)
}
