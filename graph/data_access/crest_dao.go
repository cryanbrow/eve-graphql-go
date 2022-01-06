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
		return orders, nil
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

	println(strconv.Itoa(pages))
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

	println(len(orders))
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

	crest_url, err := url.Parse(fmt.Sprintf("https://esi.evetech.net/latest/universe/systems/%s/?datasource=tranquility&language=en", strconv.Itoa(*id)))
	if err != nil {
		return system, nil
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

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}
