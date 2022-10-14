package initialize

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"order_service/global"
)

func InitTracer() (opentracing.Tracer, io.Closer) {
	jaegerInfo := global.ServiceConfig.JaegerInfo
	jaegerURL := fmt.Sprintf("http://%s:%d/api/traces", jaegerInfo.Host, jaegerInfo.Port)
	cfg := &config.Configuration{
		ServiceName: global.ServiceConfig.Name,
		Sampler:     &config.SamplerConfig{Type: jaeger.SamplerTypeConst, Param: 1},
		Reporter:    &config.ReporterConfig{LogSpans: true, CollectorEndpoint: jaegerURL},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	return tracer, closer
}
