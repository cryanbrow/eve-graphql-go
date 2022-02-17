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

const ancestryRedisKey string = "AncestryByID:"

// AncestryByID takes a context for tracing and an ID to query the Ancestry by.
func AncestryByID(ctx context.Context, id *int) (*model.Ancestry, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "AncestryByID")
	defer span.End()
	var ancestry *model.Ancestry = new(model.Ancestry)
	var err error
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}
	span.SetAttributes(attribute.Int("request.id", *id))

	inCache, result := CachingClient.CheckCache(newCtx, ancestryRedisKey+strconv.Itoa(*id))
	if !inCache {
		ancestry, err = ancestryByArray(newCtx, id)
		if err != nil {
			return nil, err
		}
		return ancestry, nil
	}
	if err := json.Unmarshal(result, &ancestry); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return ancestry, err
	}
	return ancestry, nil
}

func ancestryByArray(ctx context.Context, id *int) (*model.Ancestry, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "ancestryByArray")
	defer span.End()
	var ancestries []*model.Ancestry = make([]*model.Ancestry, 0)
	var returnAncestry *model.Ancestry
	var redisKey = ancestryRedisKey + strconv.Itoa(*id)
	baseURL := fmt.Sprintf("%s/universe/ancestries/", configuration.AppConfig.Esi.URL)

	var buffer bytes.Buffer
	responseBytes, headers, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &ancestries); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return nil, err
	}
	for _, ancestry := range ancestries {
		log.Info(*ancestry.Name)
		if *ancestry.ID == *id {
			returnAncestry = ancestry
			log.Info("Found Ancestry ID")
		}
		ancestryBytes, err := json.Marshal(*ancestry)
		if err == nil {
			CachingClient.AddToCache(newCtx, ancestryRedisKey+strconv.Itoa(*ancestry.ID), ancestryBytes, helpers.EsiTTLToMillis(newCtx, headers.Get("expires")))
		} else {
			log.Errorf(helpers.FailureMarshaling, err)
		}
	}
	span.SetAttributes(attribute.Int("request.id", *id))
	return returnAncestry, nil
}
