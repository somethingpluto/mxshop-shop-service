package main

import (
	"flag"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"goods_service/global"
	"goods_service/initialize"
	"goods_service/mode"
	"goods_service/utils/otgrpc"
	"google.golang.org/grpc"
	"io"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址:服务启动ip地址")
	Port := flag.Int("port", 8000, "port端口号: 服务启动端口号")
	Mode := flag.String("mode", "release", "mode启动模式:debug 本地调试/release 服务注册")
	flag.Parse()
	global.Port = *Port
	// 初始化文件路径
	initialize.InitFileAbsPath()
	// 初始化配置文件
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库
	initialize.InitDB()
	// 初始化Es
	initialize.InitEs()
	// 初始化tracer
	tracer, closer := initialize.InitTracer()
	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {
			panic(err)
		}
	}(closer)

	opentracing.SetGlobalTracer(tracer)
	server := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)))
	if *Mode == "debug" {
		zap.S().Warnf("debug本地调试模式 \n")
		mode.DebugMode(server, *IP)
	} else if *Mode == "release" {
		zap.S().Warnf("online 服务注册模式 \n")
		mode.ReleaseMode(server, *IP)
	}
}
