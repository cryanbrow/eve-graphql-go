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

const raceRedisKey string = "RaceByID:"

func RaceByID(ctx context.Context, id *int) (*model.Race, error) {
	newCtx, span := otel.Tracer(tracer_name).Start(ctx, "RaceByID")
	defer span.End()
	var race *model.Race = new(model.Race)
	var err error
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	inCache, result := CachingClient.CheckCache(raceRedisKey+strconv.Itoa(*id), newCtx)
	if !inCache {
		race, err = raceByArray(newCtx, id)
		if err != nil {
			return nil, err
		}
		return race, nil
	} else {
		if err := json.Unmarshal(result, &race); err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
			return race, err
		}
		return race, nil
	}
}

func raceByArray(ctx context.Context, id *int) (*model.Race, error) {
	newCtx, span := otel.Tracer(tracer_name).Start(ctx, "raceByArray")
	defer span.End()
	var races []*model.Race = make([]*model.Race, 0)
	var returnRace *model.Race
	var headers http.Header
	baseUrl := fmt.Sprintf("%s/universe/races/", configuration.AppConfig.Esi.Default.Url)
	redisKey := raceRedisKey + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, headers, err := restHelper.MakeCachingRESTCall(newCtx, baseUrl, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &races); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return nil, err
	}
	for _, race := range races {
		log.Info(*race.Name)
		if *race.RaceID == *id {
			returnRace = race
			log.Info("Found Race ID")
		}
		raceBytes, err := json.Marshal(*race)
		if err == nil {
			CachingClient.AddToCache(raceRedisKey+strconv.Itoa(*race.RaceID), raceBytes, helpers.EsiTtlToMillis(headers.Get("expires"), newCtx), newCtx)
		} else {
			log.Errorf(helpers.FailureMarshaling, err)
		}
	}
	span.SetAttributes(attribute.Int("request.id", *id))
	return returnRace, nil
}
