package data_access

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/cryanbrow/eve-graphql-go/graph/model"
)

func OrdersForRegion(regionID *int, orderType *model.Ordertype, typeID *int) ([]*model.Order, error) {
	orders := make([]*model.Order, 0)
	crest_url, err := url.Parse(fmt.Sprintf("https://esi.evetech.net/latest/markets/%s/orders/", strconv.Itoa(*regionID)))
	if err != nil {
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
		log.Printf("Error is not nil. %v", err)
	}

	for i := 2; i <= pages; i++ {
		queryParameters.Set("page", strconv.Itoa(i))
		orderResult, pages, err := ordersForRegionREST(crest_url.String())
		if err == nil && pages > 0 {
			orders = append(orders, orderResult...)
		} else {
			log.Printf("Error is not nil. %v", err)
		}
	}

	return orders, nil
}

func ordersForRegionREST(url string) ([]*model.Order, int, error) {
	var orders []*model.Order
	var pages = 0
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Could not request orders by region. %v", err)
	}
	response, err := Client.Do(request)
	if err != nil {
		log.Printf("Could not make request. %v", err)
		return orders, 0, err
	}

	pages, err = strconv.Atoi(response.Header.Get("x-pages"))

	if err != nil {
		log.Printf("Could get pages from header. %v", err)
		return orders, 0, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response for body. %v", err)
		return orders, 0, err
	}

	if err := json.Unmarshal(responseBytes, &orders); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		return orders, 0, err
	}

	return orders, pages, nil
}

func SystemByID(id *int) (*model.System, error) {
	if id == nil {
		return nil, errors.New("nil id")
	}

	inCache, result := CheckCache("SystemByID" + strconv.Itoa(*id))
	var responseBytes []byte = result
	var system *model.System = new(model.System)
	if !inCache {
		crest_url, err := url.Parse(fmt.Sprintf("https://esi.evetech.net/latest/universe/systems/%s/", strconv.Itoa(*id)))
		if err != nil {
			return nil, err
		}

		queryParameters := crest_url.Query()
		queryParameters.Add("datasource", "tranquility")
		queryParameters.Add("language", "en")

		crest_url.RawQuery = queryParameters.Encode()

		request, err := http.NewRequest(http.MethodGet, crest_url.String(), nil)
		if err != nil {
			log.Printf("Could not request orders by region. %v", err)
		}

		response, err := Client.Do(request)
		if err != nil {
			log.Printf("Could not make request. %v", err)
			return system, err
		}

		responseBytes, err = ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("Could not read response for body. %v", err)
			return system, err
		}
		AddToCache("SystemByID"+strconv.Itoa(*id), responseBytes, time.Now().UnixMilli()+43200000)
	}

	if err := json.Unmarshal(responseBytes, &system); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		return system, err
	}

	return system, nil
}

func StationByID(id *int) (*model.Station, error) {
	if id == nil {
		return nil, errors.New("nil id")
	}
	fmt.Println("Querying for station: ", id)
	var station *model.Station = new(model.Station)

	crest_url, err := url.Parse(fmt.Sprintf("https://esi.evetech.net/latest/universe/stations/%s/", strconv.Itoa(*id)))
	if err != nil {
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("language", "en")

	crest_url.RawQuery = queryParameters.Encode()

	request, err := http.NewRequest(http.MethodGet, crest_url.String(), nil)
	if err != nil {
		log.Printf("Could not request orders by region. %v", err)
	}
	response, err := Client.Do(request)
	if err != nil {
		log.Printf("Could not make request. %v", err)
		return station, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response for body. %v", err)
		return station, err
	}

	if err := json.Unmarshal(responseBytes, &station); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		return station, err
	}

	return station, nil
}

func PlanetByID(id *int) (*model.Planet, error) {
	var planet *model.Planet = new(model.Planet)
	if id == nil {
		return nil, errors.New("nil id")
	}

	crest_url, err := url.Parse(fmt.Sprintf("https://esi.evetech.net/latest/universe/planets/%s/", strconv.Itoa(*id)))
	if err != nil {
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("language", "en")

	crest_url.RawQuery = queryParameters.Encode()

	request, err := http.NewRequest(http.MethodGet, crest_url.String(), nil)
	if err != nil {
		log.Printf("Could not request orders by region. %v", err)
	}
	response, err := Client.Do(request)
	if err != nil {
		log.Printf("Could not make request. %v", err)
		return planet, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response for body. %v", err)
		return planet, err
	}

	if err := json.Unmarshal(responseBytes, &planet); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
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

	crest_url, err := url.Parse(fmt.Sprintf("https://esi.evetech.net/latest/universe/moons/%s/", strconv.Itoa(*id)))
	if err != nil {
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("language", "en")

	crest_url.RawQuery = queryParameters.Encode()

	request, err := http.NewRequest(http.MethodGet, crest_url.String(), nil)
	if err != nil {
		log.Printf("Could not request orders by region. %v", err)
	}
	response, err := Client.Do(request)
	if err != nil {
		log.Printf("Could not make request. %v", err)
		return moon, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response for body. %v", err)
		return moon, err
	}

	if err := json.Unmarshal(responseBytes, &moon); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
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

	crest_url, err := url.Parse(fmt.Sprintf("https://esi.evetech.net/latest/universe/types/%s/", strconv.Itoa(*id)))
	if err != nil {
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("language", "en")

	crest_url.RawQuery = queryParameters.Encode()

	request, err := http.NewRequest(http.MethodGet, crest_url.String(), nil)
	if err != nil {
		log.Printf("Could not request orders by region. %v", err)
	}
	response, err := Client.Do(request)
	if err != nil {
		log.Printf("Could not make request. %v", err)
		return itemType, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response for body. %v", err)
		return itemType, err
	}

	if err := json.Unmarshal(responseBytes, &itemType); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
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

	crest_url, err := url.Parse(fmt.Sprintf("https://esi.evetech.net/latest/universe/asteroid_belts/%s/", strconv.Itoa(*id)))
	if err != nil {
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("language", "en")

	crest_url.RawQuery = queryParameters.Encode()

	request, err := http.NewRequest(http.MethodGet, crest_url.String(), nil)
	if err != nil {
		log.Printf("Could not request orders by region. %v", err)
	}
	response, err := Client.Do(request)
	if err != nil {
		log.Printf("Could not make request. %v", err)
		return asteroidBelt, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response for body. %v", err)
		return asteroidBelt, err
	}

	if err := json.Unmarshal(responseBytes, &asteroidBelt); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		return asteroidBelt, err
	}

	return asteroidBelt, nil
}

func MarketGroupByID(id *int) (*model.MarketGroup, error) {
	var marketGroup *model.MarketGroup = new(model.MarketGroup)
	if id == nil {
		return nil, errors.New("nil id")
	}

	crest_url, err := url.Parse(fmt.Sprintf("https://esi.evetech.net/latest/markets/groups/%s/", strconv.Itoa(*id)))
	if err != nil {
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("language", "en")

	crest_url.RawQuery = queryParameters.Encode()

	request, err := http.NewRequest(http.MethodGet, crest_url.String(), nil)
	if err != nil {
		log.Printf("Could not request orders by region. %v", err)
	}
	response, err := Client.Do(request)
	if err != nil {
		log.Printf("Could not make request. %v", err)
		return marketGroup, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response for body. %v", err)
		return marketGroup, err
	}

	if err := json.Unmarshal(responseBytes, &marketGroup); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		return marketGroup, err
	}

	return marketGroup, nil
}

func GroupByID(id *int) (*model.Group, error) {
	var group *model.Group = new(model.Group)
	if id == nil {
		return nil, errors.New("nil id")
	}

	crest_url, err := url.Parse(fmt.Sprintf("https://esi.evetech.net/latest/universe/groups/%s/", strconv.Itoa(*id)))
	if err != nil {
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("language", "en")

	crest_url.RawQuery = queryParameters.Encode()

	request, err := http.NewRequest(http.MethodGet, crest_url.String(), nil)
	if err != nil {
		log.Printf("Could not request orders by region. %v", err)
	}
	response, err := Client.Do(request)
	if err != nil {
		log.Printf("Could not make request. %v", err)
		return group, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response for body. %v", err)
		return group, err
	}

	if err := json.Unmarshal(responseBytes, &group); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		return group, err
	}

	return group, nil
}

func GraphicByID(id *int) (*model.Graphic, error) {
	var graphic *model.Graphic = new(model.Graphic)
	if id == nil {
		return nil, errors.New("nil id")
	}

	crest_url, err := url.Parse(fmt.Sprintf("https://esi.evetech.net/latest/universe/graphics/%s/", strconv.Itoa(*id)))
	if err != nil {
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("language", "en")

	crest_url.RawQuery = queryParameters.Encode()

	request, err := http.NewRequest(http.MethodGet, crest_url.String(), nil)
	if err != nil {
		log.Printf("Could not request orders by region. %v", err)
	}
	response, err := Client.Do(request)
	if err != nil {
		log.Printf("Could not make request. %v", err)
		return graphic, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response for body. %v", err)
		return graphic, err
	}

	if err := json.Unmarshal(responseBytes, &graphic); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		return graphic, err
	}

	return graphic, nil
}

func DogmaAttributeByID(id *int) (*model.DogmaAttributeDetail, error) {
	var dogmaAttribute *model.DogmaAttributeDetail = new(model.DogmaAttributeDetail)
	if id == nil {
		return nil, nil
	}

	crest_url, err := url.Parse(fmt.Sprintf("https://esi.evetech.net/latest/dogma/attributes/%s/", strconv.Itoa(*id)))
	if err != nil {
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("language", "en")

	crest_url.RawQuery = queryParameters.Encode()

	request, err := http.NewRequest(http.MethodGet, crest_url.String(), nil)
	if err != nil {
		log.Printf("Could not request orders by region. %v", err)
	}
	response, err := Client.Do(request)
	if err != nil {
		log.Printf("Could not make request. %v", err)
		return dogmaAttribute, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response for body. %v", err)
		return dogmaAttribute, err
	}

	if err := json.Unmarshal(responseBytes, &dogmaAttribute); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		return dogmaAttribute, err
	}

	return dogmaAttribute, nil
}

func DogmaEffectByID(id *int) (*model.DogmaEffectDetail, error) {
	var dogmaEffect *model.DogmaEffectDetail = new(model.DogmaEffectDetail)
	if id == nil {
		return nil, nil
	}

	crest_url, err := url.Parse(fmt.Sprintf("https://esi.evetech.net/latest/dogma/effects/%s/", strconv.Itoa(*id)))
	if err != nil {
		return nil, err
	}

	queryParameters := crest_url.Query()
	queryParameters.Add("datasource", "tranquility")
	queryParameters.Add("language", "en")

	crest_url.RawQuery = queryParameters.Encode()

	request, err := http.NewRequest(http.MethodGet, crest_url.String(), nil)
	if err != nil {
		log.Printf("Could not request orders by region. %v", err)
	}
	response, err := Client.Do(request)
	if err != nil {
		log.Printf("Could not make request. %v", err)
		return dogmaEffect, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response for body. %v", err)
		return dogmaEffect, err
	}

	if err := json.Unmarshal(responseBytes, &dogmaEffect); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		return dogmaEffect, err
	}

	return dogmaEffect, nil
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}
