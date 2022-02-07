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

func StargateDetails(stargates []*int) ([]*model.Stargate, error) {
	stargateDetails := make([]*model.Stargate, 0)
	for _, element := range stargates {
		stargate, err := StargateByID(element)
		if err == nil {
			stargateDetails = append(stargateDetails, stargate)
		} else {
			return nil, err
		}
	}
	return stargateDetails, nil
}

func StargateByID(id *int) (*model.Stargate, error) {
	var stargate *model.Stargate = new(model.Stargate)
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}
	baseUrl := fmt.Sprintf("%s/universe/stargates/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "StargateByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return stargate, err
	}

	if err := json.Unmarshal(responseBytes, &stargate); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return stargate, err
	}

	return stargate, nil
}
