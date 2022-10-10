package main

import (
	"flag"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"order_service/initialize"
	"order_service/mode"
	otgrpc "order_service/util/otgrpcotgrpc"
)

func main() {
	IP := flag.String("ip", "127.0.0.1", "ip地址：服务启动ip地址")
	Port := flag.Int("port", 8000, "port端口号：服务启动端口号")
	Mode := flag.String("mode", "debug", "mode启动模式：debug 本地调试/release 服务注册")
	flag.Parse()
	initialize.InitFileAbsPath()
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
	initialize.InitOtherService()

	// 初始化tracer
	cfg := &config.Configuration{
		ServiceName: "order_service",
		Sampler:     &config.SamplerConfig{Type: jaeger.SamplerTypeConst, Param: 1},
		Reporter:    &config.ReporterConfig{LogSpans: true, CollectorEndpoint: "http://120.25.255.207:14268/api/traces"},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	defer closer.Close()
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
