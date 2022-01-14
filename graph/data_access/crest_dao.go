package data_access

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

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

func ordersForRegionREST(url string) ([]*model.Order, int, error) {
	var orders []*model.Order
	var pages = 0
	responseBytes, header, err := makeRESTCall(url)
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

func SystemByID(id *int) (*model.System, error) {
	if id == nil {
		return nil, errors.New("nil id")
	}

	inCache, result := CheckRedisCache("SystemByID:" + strconv.Itoa(*id))
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

		responseBytes, _, err = makeRESTCall(crest_url.String())
		if err != nil {
			return system, err
		}
		AddToRedisCache("SystemByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &system); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return system, err
	}

	return system, nil
}

func StationByID(id *int) (*model.Station, error) {
	if id == nil {
		return nil, errors.New("nil id")
	}

	inCache, result := CheckRedisCache("StationByID:" + strconv.Itoa(*id))

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

		responseBytes, _, err = makeRESTCall(crest_url.String())
		if err != nil {
			return station, err
		}
		AddToRedisCache("StationByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &station); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return station, err
	}

	return station, nil
}

func PlanetByID(id *int) (*model.Planet, error) {
	var planet *model.Planet = new(model.Planet)
	if id == nil {
		return nil, errors.New("nil id")
	}

	inCache, result := CheckRedisCache("PlanetByID:" + strconv.Itoa(*id))
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

		responseBytes, _, err = makeRESTCall(crest_url.String())
		if err != nil {
			return planet, err
		}
		AddToRedisCache("PlanetByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &planet); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return planet, err
	}

	return planet, nil
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

	inCache, result := CheckRedisCache("MoonByID:" + strconv.Itoa(*id))
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

		responseBytes, _, err = makeRESTCall(crest_url.String())
		if err != nil {
			return moon, err
		}
		AddToRedisCache("MoonByID:"+strconv.Itoa(*id), responseBytes, 43200000)
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
	inCache, result := CheckRedisCache("ItemTypeByID:" + strconv.Itoa(*id))
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

		responseBytes, _, err = makeRESTCall(crest_url.String())
		if err != nil {
			return itemType, err
		}
		AddToRedisCache("ItemTypeByID:"+strconv.Itoa(*id), responseBytes, 43200000)
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
	inCache, result := CheckRedisCache("AsteroidBeltByID:" + strconv.Itoa(*id))
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

		responseBytes, _, err = makeRESTCall(crest_url.String())
		if err != nil {
			return asteroidBelt, err
		}
		AddToRedisCache("AsteroidBeltByID:"+strconv.Itoa(*id), responseBytes, 43200000)
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
	inCache, result := CheckRedisCache("MarketGroupByID:" + strconv.Itoa(*id))
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

		responseBytes, _, err = makeRESTCall(crest_url.String())
		if err != nil {
			return marketGroup, err
		}
		AddToRedisCache("MarketGroupByID:"+strconv.Itoa(*id), responseBytes, 43200000)
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
	inCache, result := CheckRedisCache("GroupByID:" + strconv.Itoa(*id))
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

		responseBytes, _, err = makeRESTCall(crest_url.String())
		if err != nil {
			return group, err
		}
		AddToRedisCache("GroupByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &group); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return group, err
	}

	return group, nil
}

func GraphicByID(id *int) (*model.Graphic, error) {
	var graphic *model.Graphic = new(model.Graphic)
	if id == nil {
		return nil, errors.New("nil id")
	}

	inCache, result := CheckRedisCache("GraphicByID:" + strconv.Itoa(*id))
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

		responseBytes, _, err = makeRESTCall(crest_url.String())
		if err != nil {
			return graphic, err
		}
		AddToRedisCache("GraphicByID:"+strconv.Itoa(*id), responseBytes, 43200000)
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

	inCache, result := CheckRedisCache("DogmaAttributeByID:" + strconv.Itoa(*id))
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

		responseBytes, _, err = makeRESTCall(crest_url.String())
		if err != nil {
			return dogmaAttribute, err
		}
		AddToRedisCache("DogmaAttributeByID:"+strconv.Itoa(*id), responseBytes, 43200000)
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

	inCache, result := CheckRedisCache("DogmaEffectByID:" + strconv.Itoa(*id))
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

		responseBytes, _, err = makeRESTCall(crest_url.String())
		if err != nil {
			return dogmaEffect, err
		}
		AddToRedisCache("DogmaEffectByID:"+strconv.Itoa(*id), responseBytes, 43200000)
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

	inCache, result := CheckRedisCache("CategoryByID:" + strconv.Itoa(*id))
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

		responseBytes, _, err = makeRESTCall(crest_url.String())
		if err != nil {
			return category, err
		}
		AddToRedisCache("CategoryByID:"+strconv.Itoa(*id), responseBytes, 43200000)
	}

	if err := json.Unmarshal(responseBytes, &category); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return category, err
	}

	return category, nil
}

func makeRESTCall(url string) ([]byte, http.Header, error) {
	log.WithFields(log.Fields{"url": url}).Info("Making REST Call")
	request, err := http.NewRequest(http.MethodGet, url, nil)
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
