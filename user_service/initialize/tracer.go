package initialize

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"user_service/global"
)

func InitTracer() (opentracing.Tracer, io.Closer) {
	//jaegerInfo := global.ServiceConfig.JaegerInfo
	//jaegerURL := fmt.Sprintf("http://%s:%d/api/traces", jaegerInfo.Host, jaegerInfo.Port)
	cfg := &config.Configuration{
		ServiceName: global.ServiceConfig.Name,
		Sampler:     &config.SamplerConfig{Type: jaeger.SamplerTypeConst, Param: 1},
		Reporter:    &config.ReporterConfig{LogSpans: true, LocalAgentHostPort: "192.168.8.133:6831"},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	return tracer, closer
}
