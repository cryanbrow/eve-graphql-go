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

const factionRedisKey string = "FactionByID:"

// FactionByID takes a context for tracing and an ID to query the Faction by.
func FactionByID(ctx context.Context, id *int) (*model.Faction, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "CorporationByID")
	defer span.End()
	var faction *model.Faction = new(model.Faction)
	if id == nil {
		return nil, errors.New(helpers.NilID)
	}

	span.SetAttributes(attribute.Int("request.id", *id))
	inCache, result := CachingClient.CheckCache(newCtx, factionRedisKey+strconv.Itoa(*id))
	if !inCache {
		faction, err := factionByArray(newCtx, id)
		if err != nil {
			return nil, err
		}
		return faction, nil
	}
	if err := json.Unmarshal(result, &faction); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return faction, err
	}
	return faction, nil
}

func factionByArray(ctx context.Context, id *int) (*model.Faction, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "factionByArray")
	defer span.End()
	var factions []*model.Faction = make([]*model.Faction, 0)
	var returnFaction *model.Faction
	baseURL := fmt.Sprintf("%s/universe/factions/", configuration.AppConfig.Esi.Default.URL)
	redisKey := factionRedisKey + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, headers, err := restHelper.MakeCachingRESTCall(newCtx, baseURL, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &factions); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return nil, err
	}
	for _, faction := range factions {
		log.Info(*faction.Name)
		if *faction.FactionID == *id {
			returnFaction = faction
			log.Info("Found Faction ID")
		}
		factionBytes, err := json.Marshal(*faction)
		if err == nil {
			CachingClient.AddToCache(newCtx, factionRedisKey+strconv.Itoa(*faction.FactionID), factionBytes, helpers.EsiTTLToMillis(newCtx, headers.Get("expires")))
		} else {
			log.Errorf(helpers.FailureMarshaling, err)
		}
	}
	span.SetAttributes(attribute.Int("request.id", *id))
	return returnFaction, nil
}
