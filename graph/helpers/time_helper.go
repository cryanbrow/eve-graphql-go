package helpers

import (
	"context"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

// EsiTTLToMillis takes in a context for tracing purposes and a RFC1123 time for conversion to millis for TTL.
// This is done by taking the current time and subtracting the inbound time to get the millis until the TTL should occur.
func EsiTTLToMillis(ctx context.Context, esiTime string) int64 {
	_, span := otel.Tracer(tracerName).Start(ctx, "EsiTTLToMillis")
	defer span.End()
	log.WithFields(log.Fields{"time": esiTime}).Debugf("Parsing Time : %s", esiTime)
	t, err := time.Parse(time.RFC1123, esiTime)
	if err != nil {
		log.WithFields(log.Fields{"time": esiTime}).Errorf("Failed to Parse Time with Error : %v", err)
		return 43200000
	}

	log.Debugf("Responding with time: %s ", strconv.FormatInt(t.UTC().UnixMilli()-time.Now().UnixMilli(), 10))
	return t.UTC().UnixMilli() - time.Now().UnixMilli()
}
