package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"user_service/global"
)

// InitConsul
// @Description: 初始化consul 服务注册连接
//
func InitConsul() {
	var err error
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServiceConfig.ConsulInfo.Host, global.ServiceConfig.ConsulInfo.Port)

	global.Client, err = api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("服务注册 NewClient 失败", "err", err.Error())
		return
	}
	// 生成检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", "127.0.0.1", global.Port),
		GRPCUseTLS:                     false,
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "120s",
	}
	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServiceConfig.Name
	serviceID := fmt.Sprintf("%s", uuid.NewV4())
	fmt.Println(serviceID)
	global.ServiceID = serviceID
	registration.ID = serviceID
	registration.Port = global.Port
	registration.Tags = []string{"user", "service"}
	registration.Address = "127.0.0.1"
	registration.Check = check
	err = global.Client.Agent().ServiceRegister(registration)
	if err != nil {
		zap.S().Errorw("client.Agent().ServiceRegister 错误", "err", err.Error())
		return
	}
	zap.S().Infow("服务注册成功", "port", registration.Port, "ID", global.ServiceID)
}
