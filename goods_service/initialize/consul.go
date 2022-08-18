package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
	"go.uber.org/zap"
	"goods_service/global"
)

func InitRegisterService() {
	var err error
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServiceConfig.ConsulInfo.Host, global.ServiceConfig.ConsulInfo.Port)

	global.Client, err = api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("服务注册 NewClient失败", "err", err.Error())
		return
	}
	// 生成检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", "127.0.0.1", *global.FreePort),
		GRPCUseTLS:                     false,
		Timeout:                        "5s",
		Interval:                       "30s",
		DeregisterCriticalServiceAfter: "60s",
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
	registration.Port = *global.FreePort
	registration.Tags = []string{"goods", "service"}
	registration.Address = "127.0.0.1"
	registration.Check = check
	err = global.Client.Agent().ServiceRegister(registration)
	if err != nil {
		zap.S().Errorw("client.Agent().ServiceRegister 错误", "err", err.Error())
		return
	}
	zap.S().Infow("服务注册成功", "port", registration.Port, "ID", global.ServiceID)
}
