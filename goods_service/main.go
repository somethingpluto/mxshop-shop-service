package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"goods_service/handler"
	"goods_service/initialize"
	"goods_service/proto"
	"google.golang.org/grpc"
	"net"
)

func main() {
	IP := flag.String("ip", "127.0.0.1", "ip地址")
	Port := flag.Int("port", 8000, "port端口号")
	flag.Parse()

	// 初始化文件路径
	initialize.InitFileAbsPath()
	// 初始化配置文件
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库
	initialize.InitDB()
	// 微服务注册
	server := grpc.NewServer()
	proto.RegisterGoodsServer(server, &handler.GoodsServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		zap.S().Errorw("net.Listen错误", "err", err.Error())
		return
	}
	zap.S().Infof("服务启动成功 端口 %s:%d", *IP, *Port)
	err = server.Serve(listen)
	if err != nil {
		zap.S().Errorw("server.Server错误", "err", err.Error())
		return
	}

}
