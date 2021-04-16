package tracing

import (
	"fmt"
	"io"

	"github.com/joaofnds/foo/config"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
)

// initJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func InitTracer(service string) (opentracing.Tracer, io.Closer) {
	collectorEndpoint := config.JaegerCollectorEndpoint()

	cfg := &jaegerConfig.Configuration{
		ServiceName: service,
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			CollectorEndpoint: collectorEndpoint,
		},
	}
	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
