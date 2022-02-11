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

func AncestryByID(id *int, ctx context.Context) (*model.Ancestry, error) {
	newCtx, span := otel.Tracer(tracer_name).Start(ctx, "AncestryByID")
	defer span.End()
	var ancestry *model.Ancestry = new(model.Ancestry)
	var err error
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}
	span.SetAttributes(attribute.Int("request.id", *id))

	inCache, result := RedisClient.CheckRedisCache(ancestryRedisKey+strconv.Itoa(*id), newCtx)
	if !inCache {
		ancestry, err = ancestryByArray(id, newCtx)
		if err != nil {
			return nil, err
		} else {
			return ancestry, nil
		}
	} else {
		if err := json.Unmarshal(result, &ancestry); err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
			return ancestry, err
		} else {
			return ancestry, nil
		}
	}
}

func ancestryByArray(id *int, ctx context.Context) (*model.Ancestry, error) {
	newCtx, span := otel.Tracer(tracer_name).Start(ctx, "ancestryByArray")
	defer span.End()
	var ancestries []*model.Ancestry = make([]*model.Ancestry, 0)
	var returnAncestry *model.Ancestry
	var redisKey = ancestryRedisKey + strconv.Itoa(*id)
	baseUrl := fmt.Sprintf("%s/universe/ancestries/", configuration.AppConfig.Esi.Default.Url)

	var buffer bytes.Buffer
	responseBytes, headers, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey, newCtx)
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
			RedisClient.AddToRedisCache(ancestryRedisKey+strconv.Itoa(*ancestry.ID), ancestryBytes, helpers.EsiTtlToMillis(headers.Get("expires"), newCtx), newCtx)
		} else {
			log.Errorf(helpers.FailureMarshaling, err)
		}
	}
	span.SetAttributes(attribute.Int("request.id", *id))
	return returnAncestry, nil
}
