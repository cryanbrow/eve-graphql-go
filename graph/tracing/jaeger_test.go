package tracing

import (
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

func TestSuccessfulTracerProvider(t *testing.T) {
	tp, err := tracerProvider("http://localhost:11111/api/traces")
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	if tp == nil {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulSetupTracing(t *testing.T) {
	SetupTracing()
	if TraceProvider == nil {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulJaegerSetupTracing(t *testing.T) {
	configuration.AppConfig.Jaeger.Enabled = true
	SetupTracing()
	if TraceProvider == nil {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}
