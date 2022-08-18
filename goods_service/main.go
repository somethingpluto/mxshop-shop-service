package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"goods_service/global"
	"goods_service/handler"
	"goods_service/initialize"
	"goods_service/proto"
	"goods_service/util"
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
	global.FreePort = flag.Int("port", 8000, "port端口号")
	flag.Parse()

	// 初始化文件路径
	initialize.InitFileAbsPath()
	// 初始化配置文件
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库
	initialize.InitDB()

	// 判断模式
	if global.ServiceConfig.Mode != "debug" {
		freePort, err := util.GetFreePort()
		if err != nil {
			zap.S().Errorw("获取端口失败", "err", err.Error())
			return
		}
		global.FreePort = &freePort
		zap.S().Infof("获取 系统空闲端口 %d", *global.FreePort)
	}
	// 微服务注册
	server := grpc.NewServer()
	proto.RegisterGoodsServer(server, &handler.GoodsServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *global.FreePort))
	if err != nil {
		zap.S().Errorw("net.Listen错误", "err", err.Error())
		return
	}

	zap.S().Infof("服务启动成功 端口 %s:%d", *IP, *global.FreePort)
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	initialize.InitRegisterService()
	go func() {
		err = server.Serve(listen)
		if err != nil {
			panic(err)
		}
	}()

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
