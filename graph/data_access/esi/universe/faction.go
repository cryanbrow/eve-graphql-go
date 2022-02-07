package universe

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	model "github.com/cryanbrow/eve-graphql-go/graph/generated/model"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	log "github.com/sirupsen/logrus"
)

const factionRedisKey string = "FactionByID:"

func FactionByID(id *int) (*model.Faction, error) {
	var faction *model.Faction = new(model.Faction)
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}

	inCache, result := RedisClient.CheckRedisCache(factionRedisKey + strconv.Itoa(*id))
	if !inCache {
		faction, err := factionByArray(id)
		if err != nil {
			return nil, err
		} else {
			return faction, nil
		}
	} else {
		if err := json.Unmarshal(result, &faction); err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
			return faction, err
		} else {
			return faction, nil
		}
	}
}

func factionByArray(id *int) (*model.Faction, error) {
	var factions []*model.Faction = make([]*model.Faction, 0)
	var returnFaction *model.Faction
	baseUrl := fmt.Sprintf("%s/universe/factions/", configuration.AppConfig.Esi.Default.Url)
	redisKey := factionRedisKey + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, headers, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey)
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
			RedisClient.AddToRedisCache(factionRedisKey+strconv.Itoa(*faction.FactionID), factionBytes, helpers.EsiTtlToMillis(headers.Get("expires")))
		} else {
			log.Errorf(helpers.FailureMarshaling, err)
		}
	}
	return returnFaction, nil
}
