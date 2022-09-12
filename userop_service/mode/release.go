package mode

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"userop_service/global"
	"userop_service/handler"
	"userop_service/initialize"
	"userop_service/proto"
	"userop_service/util"

	"net"
	"os"
	"os/signal"
	"syscall"
)

func ReleaseMode(server *grpc.Server, ip string) {
	freePort, err := util.GetFreePort()
	if err != nil {
		zap.S().Errorw("获取端口失败", "err", err.Error())
		return
	}
	global.FreePort = freePort
	zap.S().Infof("获取 系统空闲端口 %d", global.FreePort)
	proto.RegisterMessageServer(server, &handler.UserOpService{})
	proto.RegisterAddressServer(server, &handler.UserOpService{})
	proto.RegisterUserFavoriteServer(server, &handler.UserOpService{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, global.FreePort))
	if err != nil {
		zap.S().Errorw("net.Listen错误", "err", err.Error())
		return
	}
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	initialize.InitConsul()
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
