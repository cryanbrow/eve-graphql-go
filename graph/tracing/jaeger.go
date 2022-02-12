package tracing

import (
	"log"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	var samplePercentage float64 = float64(configuration.AppConfig.Jaeger.Sample.Percent) / float64(100)
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(configuration.AppConfig.Application.Name),
			attribute.String("environment", configuration.AppConfig.Application.Environment),
		)),
		tracesdk.WithSampler(tracesdk.TraceIDRatioBased(samplePercentage)),
	)

	return tp, nil
}

func SetupTracing() {
	if configuration.AppConfig.Jaeger.Enabled {
		var err error
		//"http://localhost:14268/api/traces"
		url := configuration.AppConfig.Jaeger.Protocol + "://" + configuration.AppConfig.Jaeger.Hostname + ":" + configuration.AppConfig.Jaeger.Port + "/" + configuration.AppConfig.Jaeger.Route
		TraceProvider, err = tracerProvider(url)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		TraceProvider = tracesdk.NewTracerProvider()
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(TraceProvider)
}

var (
	TraceProvider *tracesdk.TracerProvider
)
