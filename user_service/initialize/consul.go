package initialize

import (
	"Shop_service/user_service/global"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

func InitRegisterService() {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServiceConfig.ConsulInfo.Host, global.ServiceConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("服务注册 NewClient 失败", "err", err.Error())
		return
	}
	// 生成检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("127.0.0.1:8000"),
		GRPCUseTLS:                     false,
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "120s",
	}
	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServiceConfig.Name
	registration.ID = global.ServiceConfig.Name
	registration.Port = 8000
	registration.Tags = []string{"user", "service"}
	registration.Address = "127.0.0.1"
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		zap.S().Errorw("client.Agent().ServiceRegister 错误", "err", err.Error())
		return
	}
	zap.S().Infow("服务注册成功")
}
