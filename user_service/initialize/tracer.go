package initialize

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"user_service/global"
)

func InitTracer() (opentracing.Tracer, io.Closer) {
	jaegerInfo := global.ServiceConfig.JaegerInfo
	jaegerURL := fmt.Sprintf("%s:%d", jaegerInfo.Host, jaegerInfo.Port)
	cfg := &config.Configuration{
		ServiceName: global.ServiceConfig.Name,
		Sampler:     &config.SamplerConfig{Type: jaeger.SamplerTypeConst, Param: 1},
		Reporter:    &config.ReporterConfig{LogSpans: true, LocalAgentHostPort: jaegerURL},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	return tracer, closer
}
