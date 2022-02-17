package universe

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	model "github.com/cryanbrow/eve-graphql-go/graph/generated/model"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// StarByID takes a context for tracing and an ID to query the Star by.
func StarByID(ctx context.Context, id *int) (*model.Star, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "StarByID")
	defer span.End()
	var star *model.Star = new(model.Star)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/universe/stars/%s/", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := "StarByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return star, err
	}

	if err := json.Unmarshal(responseBytes, &star); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return star, err
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	return star, nil
}
