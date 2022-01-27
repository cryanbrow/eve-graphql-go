package universe

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	log "github.com/sirupsen/logrus"

	"github.com/cryanbrow/eve-graphql-go/graph/model"
)

func AsteroidBeltDetails(asteroidBelts []*int) ([]*model.AsteroidBelt, error) {
	asteroidBeltDetails := make([]*model.AsteroidBelt, 0)
	for _, element := range asteroidBelts {
		asteroidBelt, err := AsteroidBeltByID(element)
		if err == nil {
			asteroidBeltDetails = append(asteroidBeltDetails, asteroidBelt)
		} else {
			return nil, err
		}
	}
	log.Debug(len(asteroidBeltDetails))
	return asteroidBeltDetails, nil
}

func AsteroidBeltByID(id *int) (*model.AsteroidBelt, error) {
	var asteroidBelt *model.AsteroidBelt = new(model.AsteroidBelt)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/asteroid_belts/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redis_key := "AsteroidBeltByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := helpers.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return asteroidBelt, err
	}
	log.Debug(string(responseBytes))

	if err := json.Unmarshal(responseBytes, &asteroidBelt); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return asteroidBelt, err
	}
	log.Debug(*asteroidBelt.Name)

	return asteroidBelt, nil
}

func ConstellationsByIDs(ids []*int) ([]*model.Constellation, error) {
	constellationDetails := make([]*model.Constellation, 0)
	for _, element := range ids {
		constellation, err := ConstellationByID(element)
		if err == nil {
			constellationDetails = append(constellationDetails, constellation)
		} else {
			return nil, err
		}
	}
	return constellationDetails, nil
}

func ConstellationByID(id *int) (*model.Constellation, error) {
	var constellation *model.Constellation = new(model.Constellation)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/constellations/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redis_key := "ConstellationByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := helpers.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return constellation, err
	}

	if err := json.Unmarshal(responseBytes, &constellation); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return constellation, err
	}

	return constellation, nil
}

func GroupByID(id *int) (*model.Group, error) {
	var group *model.Group = new(model.Group)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/groups/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redis_key := "GroupByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := helpers.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return group, err
	}

	if err := json.Unmarshal(responseBytes, &group); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return group, err
	}

	return group, nil
}

func ItemTypesByIDs(itemTypes []*int) ([]*model.ItemType, error) {
	itemTypeDetails := make([]*model.ItemType, 0)
	for _, element := range itemTypes {
		itemType, err := ItemTypeByID(element)
		if err == nil {
			itemTypeDetails = append(itemTypeDetails, itemType)
		} else {
			return nil, err
		}
	}
	return itemTypeDetails, nil
}

func ItemTypeByID(id *int) (*model.ItemType, error) {
	var itemType *model.ItemType = new(model.ItemType)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/types/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redis_key := "ItemTypeByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := helpers.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return itemType, err
	}

	if err := json.Unmarshal(responseBytes, &itemType); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return itemType, err
	}

	return itemType, nil
}

func MoonDetails(moons []*int) ([]*model.Moon, error) {
	moonDetails := make([]*model.Moon, 0)
	for _, element := range moons {
		moon, err := MoonByID(element)
		if err == nil {
			moonDetails = append(moonDetails, moon)
		} else {
			return nil, err
		}
	}
	return moonDetails, nil
}

func MoonByID(id *int) (*model.Moon, error) {
	var moon *model.Moon = new(model.Moon)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/moons/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redis_key := "MoonByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := helpers.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return moon, err
	}

	if err := json.Unmarshal(responseBytes, &moon); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return moon, err
	}

	return moon, nil
}

func PlanetByID(id *int) (*model.Planet, error) {
	var planet *model.Planet = new(model.Planet)
	if id == nil {
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/planets/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redis_key := "PlanetByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := helpers.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return planet, err
	}

	if err := json.Unmarshal(responseBytes, &planet); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return planet, err
	}

	return planet, nil
}

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
		return nil, errors.New("nil id")
	}
	base_url := fmt.Sprintf("%s/universe/stargates/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redis_key := "StargateByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := helpers.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return stargate, err
	}

	if err := json.Unmarshal(responseBytes, &stargate); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return stargate, err
	}

	return stargate, nil
}

func StationByArray(ids []*int) ([]*model.Station, error) {
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
		return nil, errors.New("nil id")
	}
	if *id > 2147483647 {
		return nil, nil
	}
	var station *model.Station = new(model.Station)
	base_url := fmt.Sprintf("%s/universe/stations/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redis_key := "StationByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := helpers.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return station, err
	}

	if err := json.Unmarshal(responseBytes, &station); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return station, err
	}

	return station, nil
}

func SystemByArray(ids []*int) ([]*model.System, error) {
	systemDetails := make([]*model.System, 0)
	for _, element := range ids {
		system, err := SystemByID(element)
		if err == nil {
			systemDetails = append(systemDetails, system)
		} else {
			return nil, err
		}
	}
	return systemDetails, nil
}

func SystemByID(id *int) (*model.System, error) {
	if id == nil {
		return nil, errors.New("nil id")
	}
	var system *model.System = new(model.System)
	base_url := fmt.Sprintf("%s/universe/systems/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redis_key := "SystemByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := helpers.MakeCachingRESTCall(base_url, http.MethodGet, buffer, nil, redis_key)
	if err != nil {
		return system, err
	}

	if err := json.Unmarshal(responseBytes, &system); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return system, err
	}

	return system, nil
}
