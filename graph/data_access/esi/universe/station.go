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

func StationsByIDs(ids []*int) ([]*model.Station, error) {
	stationDetails := make([]*model.Station, 0)
	for _, element := range ids {
		station, err := StationByID(element)
		if err == nil {
			stationDetails = append(stationDetails, station)
		} else {
			return nil, err
		}
	}
	return stationDetails, nil
}

func StationByID(id *int) (*model.Station, error) {
	if id == nil {
		return nil, errors.New(helpers.NilId)
	}
	if *id > 2147483647 {
		return nil, nil
	}
	var station *model.Station = new(model.Station)
	baseUrl := fmt.Sprintf("%s/universe/stations/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "StationByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return station, err
	}

	if err := json.Unmarshal(responseBytes, &station); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf(helpers.CouldNotUnmarshalResponseBytes, err)
		return station, err
	}

	return station, nil
}
