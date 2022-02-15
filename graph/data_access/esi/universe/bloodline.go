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

const bloodlineRedisKey string = "BloodlineByID:"

// BloodlineByID takes a context for tracing and an ID to query the Bloodline by.
func BloodlineByID(ctx context.Context, id *int) (*model.Bloodline, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "BloodlineByID")
	defer span.End()
	var bloodline *model.Bloodline = new(model.Bloodline)
	var err error
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	inCache, result := CachingClient.CheckCache(newCtx, bloodlineRedisKey+strconv.Itoa(*id))
	if !inCache {
		bloodline, err = bloodlineByArray(newCtx, id)
		if err != nil {
			return nil, err
		}
		return bloodline, nil
	}
	if err := json.Unmarshal(result, &bloodline); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return bloodline, err
	}
	return bloodline, nil
}

func bloodlineByArray(ctx context.Context, id *int) (*model.Bloodline, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "bloodlineByArray")
	defer span.End()
	var bloodlines []*model.Bloodline = make([]*model.Bloodline, 0)
	var returnBloodline *model.Bloodline
	baseURL := fmt.Sprintf("%s/universe/bloodlines/", configuration.AppConfig.Esi.Default.URL)
	redisKey := bloodlineRedisKey

	var buffer bytes.Buffer
	responseBytes, headers, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &bloodlines); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return nil, err
	}
	for _, bloodline := range bloodlines {
		log.Info(*bloodline.Name)
		if *bloodline.BloodlineID == *id {
			returnBloodline = bloodline
			log.Info("Found Bloodline ID")
		}
		bloodlineBytes, err := json.Marshal(*bloodline)
		if err == nil {
			CachingClient.AddToCache(newCtx, bloodlineRedisKey+strconv.Itoa(*bloodline.BloodlineID), bloodlineBytes, helpers.EsiTTLToMillis(newCtx, headers.Get("expires")))
		} else {
			log.Errorf(helpers.FailureMarshaling, err)
		}
	}
	span.SetAttributes(attribute.Int("request.id", *id))
	return returnBloodline, nil
}
