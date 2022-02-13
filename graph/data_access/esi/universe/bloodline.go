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

func BloodlineByID(ctx context.Context, id *int) (*model.Bloodline, error) {
	newCtx, span := otel.Tracer(tracer_name).Start(ctx, "BloodlineByID")
	defer span.End()
	var bloodline *model.Bloodline = new(model.Bloodline)
	var err error
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	inCache, result := CachingClient.CheckCache(bloodlineRedisKey+strconv.Itoa(*id), newCtx)
	if !inCache {
		bloodline, err = bloodlineByArray(newCtx, id)
		if err != nil {
			return nil, err
		} else {
			return bloodline, nil
		}
	} else {
		if err := json.Unmarshal(result, &bloodline); err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
			return bloodline, err
		} else {
			return bloodline, nil
		}
	}
}

func bloodlineByArray(ctx context.Context, id *int) (*model.Bloodline, error) {
	newCtx, span := otel.Tracer(tracer_name).Start(ctx, "bloodlineByArray")
	defer span.End()
	var bloodlines []*model.Bloodline = make([]*model.Bloodline, 0)
	var returnBloodline *model.Bloodline
	baseUrl := fmt.Sprintf("%s/universe/bloodlines/", configuration.AppConfig.Esi.Default.Url)
	redisKey := bloodlineRedisKey

	var buffer bytes.Buffer
	responseBytes, headers, err := restHelper.MakeCachingRESTCall(newCtx, baseUrl, http.MethodGet, buffer, nil, redisKey)
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
			CachingClient.AddToCache(bloodlineRedisKey+strconv.Itoa(*bloodline.BloodlineID), bloodlineBytes, helpers.EsiTtlToMillis(headers.Get("expires"), newCtx), newCtx)
		} else {
			log.Errorf(helpers.FailureMarshaling, err)
		}
	}
	span.SetAttributes(attribute.Int("request.id", *id))
	return returnBloodline, nil
}
