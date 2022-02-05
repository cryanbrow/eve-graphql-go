package market

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/universe"
	model "github.com/cryanbrow/eve-graphql-go/graph/generated/model"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	local_model "github.com/cryanbrow/eve-graphql-go/graph/model"
	log "github.com/sirupsen/logrus"
)

func MarketGroupByID(id *int) (*model.MarketGroup, error) {
	var marketGroup *model.MarketGroup = new(model.MarketGroup)
	if id == nil {
		return nil, errors.New("nil id")
	}
	baseUrl := fmt.Sprintf("%s/markets/groups/%s/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*id))
	redisKey := "MarketGroupByID:" + strconv.Itoa(*id)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return marketGroup, err
	}

	if err := json.Unmarshal(responseBytes, &marketGroup); err != nil {
		log.WithFields(log.Fields{"id": id}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return marketGroup, err
	}

	return marketGroup, nil
}

func OrdersForRegion(regionID *int, orderType *model.Ordertype, typeID *int, page *int) (*model.OrderWrapper, error) {
	log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID, "orderType": orderType}).Info("OrdersForRegion Called")
	orderList := make([]*model.Order, 0)
	baseUrl := fmt.Sprintf("%s/markets/%s/orders/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*regionID))

	redisKey := "OrdersForRegion:" + strconv.Itoa(*regionID) + ":" + orderType.String()

	query_params := make([]configuration.Key_value, 2)
	kv := new(configuration.Key_value)
	kv.Key = "page"
	kv.Value = strconv.Itoa(*page)
	query_params = append(query_params, *kv)

	if typeID != nil {
		redisKey = redisKey + ":" + strconv.Itoa(*typeID)
		kv.Key = "type_id"
		kv.Value = strconv.Itoa(*typeID)
		query_params = append(query_params, *kv)
	}

	redisKey = redisKey + ":" + strconv.Itoa(*page)

	orderResult, pages, err := ordersForRegionREST(baseUrl, query_params, redisKey)

	if err == nil {
		orderList = append(orderList, orderResult...)
	} else {
		log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID, "orderType": orderType}).Errorf("First page query for Orders has error : %v", err)
		println(err)
	}

	return_orders := new(model.OrderWrapper)
	return_orders.List = orderList
	return_orders.Xpages = &pages

	return return_orders, nil
}

func OrdersForRegionByName(region *string, orderType *model.Ordertype, typeName *string, page *int) (*model.OrderWrapper, error) {
	regionID, err := universe.IdForName(region, local_model.REGIONS)
	if err != nil {
		return nil, errors.New("unknown name for region")
	}
	typeID, err := universe.IdForName(typeName, local_model.INVENTORY_TYPES)
	if err != nil {
		return nil, errors.New("unknown name for typeName")
	}
	orders, err := OrdersForRegion(&regionID, orderType, &typeID, page)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func ordersForRegionREST(url string, additional_query_params []configuration.Key_value, redisKey string) ([]*model.Order, int, error) {
	var orders []*model.Order
	var pages = 0
	var buffer bytes.Buffer
	responseBytes, header, err := restHelper.MakeCachingRESTCall(url, http.MethodGet, buffer, additional_query_params, redisKey)
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

func OrderHistory(regionID *int, typeID *int) ([]*model.OrderHistory, error) {
	if regionID == nil || typeID == nil {
		return nil, errors.New("nil id")
	}
	var orderHistory []*model.OrderHistory = make([]*model.OrderHistory, 0)
	baseUrl := fmt.Sprintf("%s/markets/%s/history", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*regionID))
	redisKey := "OrderHistoryByID:" + strconv.Itoa(*regionID) + ":" + strconv.Itoa(*typeID)

	var buffer bytes.Buffer
	responseBytes, _, err := restHelper.MakeCachingRESTCall(baseUrl, http.MethodGet, buffer, nil, redisKey)
	if err != nil {
		return orderHistory, err
	}

	if err := json.Unmarshal(responseBytes, &orderHistory); err != nil {
		log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID}).Errorf("Could not unmarshal reponseBytes. %v", err)
		return orderHistory, err
	}

	return orderHistory, nil
}

type RestHelper interface {
	MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error)
}

var (
	restHelper RestHelper
)

func SetupMarketRest() {
	restHelper = &helpers.RestHelperClient{}
}
