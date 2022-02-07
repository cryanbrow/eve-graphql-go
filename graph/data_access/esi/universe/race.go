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

const raceRedisKey string = "RaceByID:"

func RaceByID(id *int) (*model.Race, error) {
	var race *model.Race = new(model.Race)
	var err error
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}

	inCache, result := RedisClient.CheckRedisCache(raceRedisKey + strconv.Itoa(*id))
	if !inCache {
		race, err = raceByArray(id)
		if err != nil {
			return nil, err
		} else {
			return race, nil
		}
	} else {
		if err := json.Unmarshal(result, &race); err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
			return race, err
		} else {
			return race, nil
		}
	}
}

func raceByArray(id *int) (*model.Race, error) {
	var races []*model.Race = make([]*model.Race, 0)
	var returnRace *model.Race
	var headers http.Header = nil
	baseUrl := fmt.Sprintf("%s/universe/races/", configuration.AppConfig.Esi.Default.Url)
	redisKey := raceRedisKey + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, headers, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey)
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
			RedisClient.AddToRedisCache(raceRedisKey+strconv.Itoa(*race.RaceID), raceBytes, helpers.EsiTtlToMillis(headers.Get("expires")))
		} else {
			log.Errorf(helpers.FailureMarshaling, err)
		}
	}
	return returnRace, nil
}
