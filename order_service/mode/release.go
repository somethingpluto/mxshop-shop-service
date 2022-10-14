package mode

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"order_service/global"
	"order_service/handler"
	"order_service/proto"
	"order_service/util"

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
	proto.RegisterOrderServer(server, &handler.OrderService{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, global.FreePort))
	if err != nil {
		zap.S().Errorw("net.Listen错误", "err", err.Error())
		return
	}
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	go func() {
		err = server.Serve(listen)
		panic(err)
	}()

	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServiceConfig.ConsulInfo.Host, global.ServiceConfig.ConsulInfo.Port)
	global.Client, err = api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("服务注册 NewClient失败", "err", err.Error())
		return
	}
	// 生成检查对象
	checkInfo := global.ServiceConfig.RegisterInfo
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", global.ServiceConfig.Host, global.FreePort),
		GRPCUseTLS:                     false,
		Timeout:                        checkInfo.CheckTimeOut,
		Interval:                       checkInfo.CheckInterval,
		DeregisterCriticalServiceAfter: checkInfo.DeregisterTime,
	}
	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServiceConfig.Name
	v4, err := uuid.NewV4()
	if err != nil {
		zap.S().Errorw("uuid.NewV4 failed", "err", err.Error())
		return
	}
	serviceID := v4.String()
	global.ServiceID = serviceID
	registration.ID = serviceID
	registration.Port = global.FreePort
	registration.Tags = checkInfo.Tags
	registration.Address = global.ServiceConfig.Host
	registration.Check = check
	err = global.Client.Agent().ServiceRegister(registration)
	if err != nil {
		zap.S().Errorw("client.Agent().ServiceRegister 错误", "err", err.Error())
		return
	}
	zap.S().Infow("服务注册成功", "port", registration.Port, "ID", global.ServiceID)

	// 优雅关机
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
