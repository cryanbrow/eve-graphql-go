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

const bloodlineRedisKey string = "BloodlineByID:"

func BloodlineByID(id *int) (*model.Bloodline, error) {
	var bloodline *model.Bloodline = new(model.Bloodline)
	var err error
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}

	inCache, result := RedisClient.CheckRedisCache(bloodlineRedisKey + strconv.Itoa(*id))
	if !inCache {
		bloodline, err = bloodlineByArray(id)
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

func bloodlineByArray(id *int) (*model.Bloodline, error) {
	var bloodlines []*model.Bloodline = make([]*model.Bloodline, 0)
	var returnBloodline *model.Bloodline
	baseUrl := fmt.Sprintf("%s/universe/bloodlines/", configuration.AppConfig.Esi.Default.Url)
	redisKey := bloodlineRedisKey

	var buffer bytes.Buffer
	responseBytes, headers, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey)
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
			RedisClient.AddToRedisCache(bloodlineRedisKey+strconv.Itoa(*bloodline.BloodlineID), bloodlineBytes, helpers.EsiTtlToMillis(headers.Get("expires")))
		} else {
			log.Errorf(helpers.FailureMarshaling, err)
		}
	}
	return returnBloodline, nil
}
