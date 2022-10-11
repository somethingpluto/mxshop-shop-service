package mode

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
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
	zap.S().Infow("Info", "message", fmt.Sprintf("获取主机端口: %d", global.Port))
	proto.RegisterUserServer(server, &handler.UserService{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, global.Port))
	if err != nil {
		zap.S().Errorw("net.Listen错误", "err", err.Error())
		return
	}
	registerService()
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

func registerService() {
	var err error
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServiceConfig.ConsulInfo.Host, global.ServiceConfig.ConsulInfo.Port)

	global.Client, err = api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("服务注册 NewClient 失败", "err", err.Error())
		return
	}
	// 生成检查对象
	checkConfig := global.ServiceConfig.ServiceInfo
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", "192.168.8.1", global.Port),
		GRPCUseTLS:                     false,
		Timeout:                        "5s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: checkConfig.DeregisterTime,
	}
	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServiceConfig.Name
	serviceID := fmt.Sprintf("%s", uuid.NewV4())
	fmt.Println(serviceID)
	global.ServiceID = serviceID
	registration.ID = serviceID
	registration.Port = global.Port
	registration.Tags = checkConfig.Tags
	registration.Address = "192.168.8.1"
	registration.Check = check
	err = global.Client.Agent().ServiceRegister(registration)
	if err != nil {
		zap.S().Errorw("Error", "message", "client.Agent().ServiceRegister 错误", "err", err.Error())
		return
	}
	zap.S().Infow("Info", "message", "服务注册成功", "port", registration.Port, "ID", global.ServiceID)
}
