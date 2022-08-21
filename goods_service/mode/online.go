package mode

import (
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

func OnlineMode(server *grpc.Server, ip string) {
	freePort, err := util.GetFreePort()
	if err != nil {
		zap.S().Errorw("获取端口失败", "err", err.Error())
		return
	}
	global.FreePort = freePort
	zap.S().Infof("获取 系统空闲端口 %d", global.FreePort)
	proto.RegisterGoodsServer(server, &handler.GoodsServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, global.FreePort))
	if err != nil {
		zap.S().Errorw("net.Listen错误", "err", err.Error())
		return
	}
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	initialize.InitRegisterService()
	go func() {
		err = server.Serve(listen)
		panic(err)
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
