package main

import (
	"flag"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
	"userop_service/initialize"
	"userop_service/mode"
	"userop_service/util/otgrpc"
)

func main() {
	// 命令行参数
	IP := flag.String("ip", "0.0.0.0", "ip地址：服务启动ip地址")
	Port := flag.Int("port", 8000, "port端口号：服务启动端口号")
	Mode := flag.String("mode", "release", "mode启动模式：debug 本地调试/release 服务注册")
	flag.Parse()

	initialize.InitFileAbsPath()
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
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
		mode.DebugMode(server, *IP, *Port)
	} else if *Mode == "release" {
		zap.S().Warnf("release服务注册模式 \n")
		mode.ReleaseMode(server, *IP)
	}
}
