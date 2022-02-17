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

const asteroidBeltRedisKey string = "AsteroidBeltByID:"

// AsteroidBeltDetails takes in a context for tracing and an array of astroidBelt IDs to query buy to return the data requested.
func AsteroidBeltDetails(ctx context.Context, asteroidBelts []*int) ([]*model.AsteroidBelt, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "AsteroidBeltDetails")
	defer span.End()
	asteroidBeltDetails := make([]*model.AsteroidBelt, 0)
	for _, element := range asteroidBelts {
		asteroidBelt, err := AsteroidBeltByID(newCtx, element)
		if err == nil {
			asteroidBeltDetails = append(asteroidBeltDetails, asteroidBelt)
		} else {
			return nil, err
		}
	}
	log.Debug(len(asteroidBeltDetails))
	return asteroidBeltDetails, nil
}

// AsteroidBeltByID takes a context for tracing and an ID to query the Astroid Belt by.
func AsteroidBeltByID(ctx context.Context, id *int) (*model.AsteroidBelt, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "AsteroidBeltByID")
	defer span.End()
	var asteroidBelt *model.AsteroidBelt = new(model.AsteroidBelt)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	baseURL := fmt.Sprintf("%s/universe/asteroid_belts/%s/", configuration.AppConfig.Esi.URL, strconv.Itoa(*id))
	redisKey := asteroidBeltRedisKey + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return asteroidBelt, err
	}
	log.Debug(string(responseBytes))

	if err := json.Unmarshal(responseBytes, &asteroidBelt); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return asteroidBelt, err
	}
	log.Debug(*asteroidBelt.Name)

	span.SetAttributes(attribute.Int("request.id", *id))
	return asteroidBelt, nil
}
