package tracing

import (
	"time"

	"github.com/Hendra-Huang/go-standard-layout/log"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

// NewTracer creates a new instance of Jaeger tracer.
func NewTracer(serviceName string, metricsFactory metrics.Factory, hostPort string) opentracing.Tracer {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  hostPort,
		},
	}
	jLogger := jaegerlog.NullLogger // use StdLogger if need to debug
	tracer, _, err := cfg.New(
		serviceName,
		config.Logger(jLogger),
		config.Metrics(metricsFactory),
	)
	if err != nil {
		log.Fatal("cannot initialize Jaeger Tracer")
	}

	return tracer
}
