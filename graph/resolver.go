package graph

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

//go:generate go get github.com/99designs/gqlgen/cmd@v0.14.0
//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

const tracer_name = "github.com/cryanbrow/eve-graphql-go/graph"

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(configuration.AppConfig.Application.Name),
			attribute.String("environment", configuration.AppConfig.Application.Environment),
		)),
		tracesdk.WithSampler(tracesdk.TraceIDRatioBased(configuration.AppConfig.Jaeger.Sample.Percent)),
	)
	return tp, nil
}

func SetupResolver() {
	var err error = nil
	traceProvider, err = tracerProvider(configuration.AppConfig.Jaeger.Protocol + "://" + configuration.AppConfig.Jaeger.Hostname + ":" + configuration.AppConfig.Jaeger.Port + "/" + configuration.AppConfig.Jaeger.Route)
	//"http://localhost:14268/api/traces"
	if err != nil {
		log.Fatal(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(traceProvider)
}

var (
	traceProvider *tracesdk.TracerProvider
)
