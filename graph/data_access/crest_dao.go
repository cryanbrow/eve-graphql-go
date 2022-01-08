package data_access

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

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
	var system *model.System = new(model.System)

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

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response for body. %v", err)
		return system, err
	}

	if err := json.Unmarshal(responseBytes, &system); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		return system, err
	}

	return system, nil
}

func StationByID(id *int) (*model.Station, error) {
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

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}
