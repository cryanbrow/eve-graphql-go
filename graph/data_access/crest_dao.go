package data_access

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	cache "github.com/cryanbrow/eve-graphql-go/graph/caching"
	log "github.com/sirupsen/logrus"

	"github.com/cryanbrow/eve-graphql-go/graph/model"
)

func OrdersForRegion(regionID *int, orderType *model.Ordertype, typeID *int) ([]*model.Order, error) {
	log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID, "orderType": orderType}).Info("OrdersForRegion Called")
	orders := make([]*model.Order, 0)
	crest_url, err := url.Parse(fmt.Sprintf("%s/markets/%s/orders/", baseUriESI, strconv.Itoa(*regionID)))
	if err != nil {
		log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID, "orderType": orderType}).Errorf("Failed to Parse URL with Error : %s", err)
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("order_type", string(*orderType))
	queryParameters.Add("page", "1")
	if typeID != nil {
		queryParameters.Add("type_id", strconv.Itoa(*typeID))
	}
	crest_url.RawQuery = queryParameters.Encode()

	orderResult, pages, err := ordersForRegionREST(crest_url.String())

	if err == nil && pages > 0 {
		orders = append(orders, orderResult...)
	} else {
		log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID, "orderType": orderType}).Errorf("First page query for Orders has error : %v", err)
		println(err)
	}

	for i := 2; i <= pages; i++ {
		queryParameters.Set("page", strconv.Itoa(i))
		orderResult, pages, err := ordersForRegionREST(crest_url.String())
		if err == nil && pages > 0 {
			orders = append(orders, orderResult...)
		} else {
			log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID, "orderType": orderType, "page": i}).Errorf("Error is not nil : %v", err)
		}
	}

	return orders, nil
}

func OrdersForRegionByName(region *string, orderType *model.Ordertype, typeID *string) ([]*model.Order, error) {
	return nil, nil
}

func OrderHistory(regionID *int, typeID *int) ([]*model.OrderHistory, error) {
	if regionID == nil || typeID == nil {
		return nil, errors.New("nil id")
	}

	inCache, result := cache.CheckRedisCache("OrderHistoryByID:" + strconv.Itoa(*regionID) + ":" + strconv.Itoa(*typeID))
	var responseBytes []byte = result
	var orderHistory []*model.OrderHistory = make([]*model.OrderHistory, 0)
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/markets/%s/history", baseUriESI, strconv.Itoa(*regionID)))
		if err != nil {
			log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("type_id", strconv.Itoa(*typeID))
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return orderHistory, err
		}
		cache.AddToRedisCache("OrderHistoryByID:"+strconv.Itoa(*regionID)+":"+strconv.Itoa(*typeID), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &orderHistory); err != nil {
		log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return orderHistory, err
	}

	return orderHistory, nil
}

func ordersForRegionREST(url string) ([]*model.Order, int, error) {
	var orders []*model.Order
	var pages = 0
	var buffer bytes.Buffer
	responseBytes, header, err := makeRESTCall(url, http.MethodGet, buffer)
	if err != nil {
		return orders, 0, err
	}

	pages, _ = strconv.Atoi(header.Get("x-pages"))

	if err := json.Unmarshal(responseBytes, &orders); err != nil {
		log.WithFields(log.Fields{"url": url}).Errorf("Could not unmarshal reponseBytes. : %v", err.Error())
		return orders, 0, err
	}

	return orders, pages, nil
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

	inCache, result := cache.CheckRedisCache("SystemByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	var system *model.System = new(model.System)
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/universe/systems/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return system, err
		}
		cache.AddToRedisCache("SystemByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &system); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return system, err
	}

	return system, nil
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

	inCache, result := cache.CheckRedisCache("StationByID:" + strconv.Itoa(*id))

	var station *model.Station = new(model.Station)
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/universe/stations/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return station, err
		}
		cache.AddToRedisCache("StationByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &station); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return station, err
	}

	return station, nil
}

func CorporationByID(id *int) (*model.Corporation, error) {
	if id == nil {
		return nil, errors.New("nil id")
	}

	inCache, result := cache.CheckRedisCache("CorporationByID:" + strconv.Itoa(*id))

	var corporation *model.Corporation = new(model.Corporation)
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/corporations/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return corporation, err
		}
		cache.AddToRedisCache("CorporationByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &corporation); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return corporation, err
	}

	return corporation, nil
}

func AllianceByID(id *int) (*model.Alliance, error) {
	if id == nil {
		return nil, errors.New("nil id")
	}

	inCache, result := cache.CheckRedisCache("AllianceByID:" + strconv.Itoa(*id))

	var alliance *model.Alliance = new(model.Alliance)
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/alliances/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return alliance, err
		}
		cache.AddToRedisCache("AllianceByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &alliance); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return alliance, err
	}

	return alliance, nil
}

func CharacterByID(id *int) (*model.Character, error) {
	if id == nil {
		return nil, errors.New("nil id")
	}

	inCache, result := cache.CheckRedisCache("CharacterByID:" + strconv.Itoa(*id))

	var character *model.Character = new(model.Character)
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/characters/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return character, err
		}
		cache.AddToRedisCache("CharacterByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &character); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return character, err
	}

	return character, nil
}

func PlanetByID(id *int) (*model.Planet, error) {
	var planet *model.Planet = new(model.Planet)
	if id == nil {
		return nil, errors.New("nil id")
	}

	inCache, result := cache.CheckRedisCache("PlanetByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	if !inCache {

		crest_url, err := url.Parse(fmt.Sprintf("%s/universe/planets/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return planet, err
		}
		cache.AddToRedisCache("PlanetByID:"+strconv.Itoa(*id), responseBytes, 43200000)
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

	inCache, result := cache.CheckRedisCache("StargateByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	if !inCache {

		crest_url, err := url.Parse(fmt.Sprintf("%s/universe/stargates/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return stargate, err
		}
		cache.AddToRedisCache("StargateByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &stargate); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return stargate, err
	}

	return stargate, nil
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

	inCache, result := cache.CheckRedisCache("MoonByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/universe/moons/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return moon, err
		}
		cache.AddToRedisCache("MoonByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &moon); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return moon, err
	}

	return moon, nil
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
	inCache, result := cache.CheckRedisCache("ItemTypeByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/universe/types/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorln("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return itemType, err
		}
		cache.AddToRedisCache("ItemTypeByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &itemType); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return itemType, err
	}

	return itemType, nil
}

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
	return asteroidBeltDetails, nil
}

func AsteroidBeltByID(id *int) (*model.AsteroidBelt, error) {
	var asteroidBelt *model.AsteroidBelt = new(model.AsteroidBelt)
	if id == nil {
		return nil, errors.New("nil id")
	}
	inCache, result := cache.CheckRedisCache("AsteroidBeltByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/universe/asteroid_belts/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return asteroidBelt, err
		}
		cache.AddToRedisCache("AsteroidBeltByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &asteroidBelt); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return asteroidBelt, err
	}

	return asteroidBelt, nil
}

func MarketGroupByID(id *int) (*model.MarketGroup, error) {
	var marketGroup *model.MarketGroup = new(model.MarketGroup)
	if id == nil {
		return nil, errors.New("nil id")
	}
	inCache, result := cache.CheckRedisCache("MarketGroupByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/markets/groups/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return marketGroup, err
		}
		cache.AddToRedisCache("MarketGroupByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &marketGroup); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return marketGroup, err
	}

	return marketGroup, nil
}

func GroupByID(id *int) (*model.Group, error) {
	var group *model.Group = new(model.Group)
	if id == nil {
		return nil, errors.New("nil id")
	}
	inCache, result := cache.CheckRedisCache("GroupByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/universe/groups/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return group, err
		}
		cache.AddToRedisCache("GroupByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &group); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return group, err
	}

	return group, nil
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
	inCache, result := cache.CheckRedisCache("ConstellationByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/universe/constellations/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return constellation, err
		}
		cache.AddToRedisCache("ConstellationByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &constellation); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return constellation, err
	}

	return constellation, nil
}

func StarByID(id *int) (*model.Star, error) {
	var star *model.Star = new(model.Star)
	if id == nil {
		return nil, errors.New("nil id")
	}
	inCache, result := cache.CheckRedisCache("StarByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/universe/stars/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return star, err
		}
		cache.AddToRedisCache("StarByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &star); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return star, err
	}

	return star, nil
}

func GraphicByID(id *int) (*model.Graphic, error) {
	var graphic *model.Graphic = new(model.Graphic)
	if id == nil {
		return nil, errors.New("nil id")
	}

	inCache, result := cache.CheckRedisCache("GraphicByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/universe/graphics/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return graphic, err
		}
		cache.AddToRedisCache("GraphicByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &graphic); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return graphic, err
	}

	return graphic, nil
}

func DogmaAttributeByID(id *int) (*model.DogmaAttributeDetail, error) {
	var dogmaAttribute *model.DogmaAttributeDetail = new(model.DogmaAttributeDetail)
	if id == nil {
		return nil, nil
	}

	inCache, result := cache.CheckRedisCache("DogmaAttributeByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/dogma/attributes/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return dogmaAttribute, err
		}
		cache.AddToRedisCache("DogmaAttributeByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &dogmaAttribute); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return dogmaAttribute, err
	}

	return dogmaAttribute, nil
}

func DogmaEffectByID(id *int) (*model.DogmaEffectDetail, error) {
	var dogmaEffect *model.DogmaEffectDetail = new(model.DogmaEffectDetail)
	if id == nil {
		return nil, nil
	}

	inCache, result := cache.CheckRedisCache("DogmaEffectByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/dogma/effects/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return dogmaEffect, err
		}
		cache.AddToRedisCache("DogmaEffectByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &dogmaEffect); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return dogmaEffect, err
	}

	return dogmaEffect, nil
}

func CategoryByID(id *int) (*model.Category, error) {
	var category *model.Category = new(model.Category)
	if id == nil {
		return nil, nil
	}

	inCache, result := cache.CheckRedisCache("CategoryByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/universe/categories/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return category, err
		}
		cache.AddToRedisCache("CategoryByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &category); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return category, err
	}

	return category, nil
}

func RegionByID(id *int) (*model.Region, error) {
	var region *model.Region = new(model.Region)
	if id == nil {
		return nil, nil
	}

	inCache, result := cache.CheckRedisCache("RegionByID:" + strconv.Itoa(*id))
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/universe/regions/%s/", baseUriESI, strconv.Itoa(*id)))
		if err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		var buffer bytes.Buffer
		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
		if err != nil {
			return region, err
		}
		cache.AddToRedisCache("RegionByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &region); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return region, err
	}

	return region, nil
}

func FactionByID(id *int) (*model.Faction, error) {
	var faction *model.Faction = new(model.Faction)
	if id == nil {
		return nil, nil
	}

	inCache, result := cache.CheckRedisCache("FactionByID:" + strconv.Itoa(*id))
	if !inCache {
		faction, err := factionByArray(id)
		if err != nil {
			return nil, err
		} else {
			return faction, nil
		}
	} else {
		if err := json.Unmarshal(result, &faction); err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
			return faction, err
		} else {
			return faction, nil
		}
	}
}

func factionByArray(id *int) (*model.Faction, error) {
	var factions []*model.Faction = make([]*model.Faction, 0)
	var returnFaction *model.Faction
	var responseBytes []byte = make([]byte, 0)
	crest_url, err := url.Parse(fmt.Sprintf("%s/universe/factions/", baseUriESI))
	if err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("language", "en")

	crest_url.RawQuery = queryParameters.Encode()

	var buffer bytes.Buffer
	responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &factions); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
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
			cache.AddToRedisCache("FactionByID:"+strconv.Itoa(*faction.FactionID), factionBytes, 43200000)
		} else {
			log.Errorf("Failure Marshalling: %v", err)
		}
	}
	return returnFaction, nil
}

func AncestryByID(id *int) (*model.Ancestry, error) {
	var ancestry *model.Ancestry = new(model.Ancestry)
	var err error
	if id == nil {
		return nil, nil
	}

	inCache, result := cache.CheckRedisCache("AncestryByID:" + strconv.Itoa(*id))
	if !inCache {
		ancestry, err = ancestryByArray(id)
		if err != nil {
			return nil, err
		} else {
			return ancestry, nil
		}
	} else {
		if err := json.Unmarshal(result, &ancestry); err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
			return ancestry, err
		} else {
			return ancestry, nil
		}
	}
}

func ancestryByArray(id *int) (*model.Ancestry, error) {
	var ancestries []*model.Ancestry = make([]*model.Ancestry, 0)
	var returnAncestry *model.Ancestry
	var responseBytes []byte = make([]byte, 0)
	crest_url, err := url.Parse(fmt.Sprintf("%s/universe/ancestries/", baseUriESI))
	if err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("language", "en")

	crest_url.RawQuery = queryParameters.Encode()

	var buffer bytes.Buffer
	responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &ancestries); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
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
			cache.AddToRedisCache("AncestryByID:"+strconv.Itoa(*ancestry.ID), ancestryBytes, 43200000)
		} else {
			log.Errorf("Failure Marshalling: %v", err)
		}
	}
	return returnAncestry, nil
}

func BloodlineByID(id *int) (*model.Bloodline, error) {
	var bloodline *model.Bloodline = new(model.Bloodline)
	var err error
	if id == nil {
		return nil, nil
	}

	inCache, result := cache.CheckRedisCache("BloodlineByID:" + strconv.Itoa(*id))
	if !inCache {
		bloodline, err = bloodlineByArray(id)
		if err != nil {
			return nil, err
		} else {
			return bloodline, nil
		}
	} else {
		if err := json.Unmarshal(result, &bloodline); err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
			return bloodline, err
		} else {
			return bloodline, nil
		}
	}
}

func bloodlineByArray(id *int) (*model.Bloodline, error) {
	var bloodlines []*model.Bloodline = make([]*model.Bloodline, 0)
	var returnBloodline *model.Bloodline
	var responseBytes []byte = make([]byte, 0)
	crest_url, err := url.Parse(fmt.Sprintf("%s/universe/bloodlines/", baseUriESI))
	if err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("language", "en")

	crest_url.RawQuery = queryParameters.Encode()

	var buffer bytes.Buffer
	responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &bloodlines); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
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
			cache.AddToRedisCache("BloodlineByID:"+strconv.Itoa(*bloodline.BloodlineID), bloodlineBytes, 43200000)
		} else {
			log.Errorf("Failure Marshalling: %v", err)
		}
	}
	return returnBloodline, nil
}

func RaceByID(id *int) (*model.Race, error) {
	var race *model.Race = new(model.Race)
	var err error
	if id == nil {
		return nil, nil
	}

	inCache, result := cache.CheckRedisCache("RaceByID:" + strconv.Itoa(*id))
	if !inCache {
		race, err = raceByArray(id)
		if err != nil {
			return nil, err
		} else {
			return race, nil
		}
	} else {
		if err := json.Unmarshal(result, &race); err != nil {
			log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
			return race, err
		} else {
			return race, nil
		}
	}
}

func raceByArray(id *int) (*model.Race, error) {
	var races []*model.Race = make([]*model.Race, 0)
	var returnRace *model.Race
	var responseBytes []byte = make([]byte, 0)
	crest_url, err := url.Parse(fmt.Sprintf("%s/universe/races/", baseUriESI))
	if err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Failed to Parse URL with Error : %v", err)
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("language", "en")

	crest_url.RawQuery = queryParameters.Encode()

	var buffer bytes.Buffer
	responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodGet, buffer)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseBytes, &races); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
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
			cache.AddToRedisCache("RaceByID:"+strconv.Itoa(*race.RaceID), raceBytes, 43200000)
		} else {
			log.Errorf("Failure Marshalling: %v", err)
		}
	}
	return returnRace, nil
}

func idForName(name *string) (int, error) {
	var ids *model.Names = new(model.Names)
	if name == nil {
		return 0, errors.New("nil name")
	}

	inCache, result := cache.CheckRedisCache("IDForName:" + *name)
	var responseBytes []byte = result
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("%s/universe/ids/", baseUriESI))
		if err != nil {
			log.WithFields(log.Fields{"name": name}).Errorf("Failed to Parse URL with Error : %v", err)
			return 0, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		singleItemArray := []string{*name}

		var buf bytes.Buffer
		err = json.NewEncoder(&buf).Encode(singleItemArray)
		if err != nil {
			log.Error(err)
			return 0, err
		}

		responseBytes, _, err = makeRESTCall(crest_url.String(), http.MethodPost, buf)
		if err != nil {
			return 0, err
		}
		cache.AddToRedisCache("IDForName:"+*name, responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &ids); err != nil {
		log.WithFields(log.Fields{"name": *name}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return 0, err
	}

	switch {
	case ids.Agents != nil:
		return *ids.Agents[0].ID, nil
	case ids.Alliances != nil:
		return *ids.Alliances[0].ID, nil
	case ids.Characters != nil:
		return *ids.Characters[0].ID, nil
	case ids.Constellations != nil:
		return *ids.Constellations[0].ID, nil
	case ids.Corporations != nil:
		return *ids.Corporations[0].ID, nil
	case ids.Factions != nil:
		return *ids.Factions[0].ID, nil
	case ids.InventoryTypes != nil:
		return *ids.InventoryTypes[0].ID, nil
	case ids.Regions != nil:
		return *ids.Regions[0].ID, nil
	case ids.Systems != nil:
		return *ids.Systems[0].ID, nil
	default:
		return 0, errors.New("all fields nil")
	}
}

func makeRESTCall(url string, verb string, byteBuffer bytes.Buffer) ([]byte, http.Header, error) {
	log.WithFields(log.Fields{"url": url}).Info("Making REST Call")
	request, err := http.NewRequest(verb, url, &byteBuffer)
	if err != nil {
		log.WithFields(log.Fields{"url": url}).Errorf("Could not build request. : %v", err)
	}
	response, err := Client.Do(request)
	if err != nil {
		log.WithFields(log.Fields{"url": url}).Errorf("Could not make request. : %v", err)
		return make([]byte, 0), nil, err
	}

	h := response.Header
	responseBytes, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		log.WithFields(log.Fields{"url": url, "status_code": response.StatusCode}).Errorf("Received bad status code. : %v", err)
		return make([]byte, 0), nil, err
	}
	if err != nil {
		log.WithFields(log.Fields{"url": url}).Errorf("Could not read response for body. : %v", err)
		return make([]byte, 0), nil, err
	}
	return responseBytes, h, nil
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	baseUriESI string
	Client     HTTPClient
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	Client = &http.Client{}
	baseUriESI = "https://esi.evetech.net/latest"
}
