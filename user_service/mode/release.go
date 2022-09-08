package mode

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"syscall"
	"user_service/global"
	"user_service/handler"
	"user_service/initialize"
	"user_service/proto"
	"user_service/util"
)

// ReleaseMode
// @Description: release服务注册模式
// @param server
// @param ip
//
func ReleaseMode(server *grpc.Server, ip string) {
	var err error
	freePort, err := util.GetFreePort()
	if err != nil {
		zap.S().Errorw("获取端口失败", "err", err.Error())
		return
	}
	global.Port = freePort
	zap.S().Infof("获取 系统空闲端口 %d", global.Port)
	proto.RegisterUserServer(server, &handler.UserService{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, global.Port))
	if err != nil {
		zap.S().Errorw("net.Listen错误", "err", err.Error())
		return
	}
	initialize.InitConsul()
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
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
