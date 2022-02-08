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
	local_model "github.com/cryanbrow/eve-graphql-go/graph/model"
	log "github.com/sirupsen/logrus"
)

func OrdersForRegion(regionID *int, orderType *model.Ordertype, typeID *int, page *int) (*model.OrderWrapper, error) {
	log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID, "orderType": orderType}).Info("OrdersForRegion Called")
	orderList := make([]*model.Order, 0)
	baseUrl := fmt.Sprintf("%s/markets/%s/orders/", configuration.AppConfig.Esi.Default.Url, strconv.Itoa(*regionID))

	redisKey := "OrdersForRegion:" + strconv.Itoa(*regionID) + ":" + orderType.String()

	queryParams := make([]configuration.Key_value, 2)
	kv := new(configuration.Key_value)
	kv.Key = "page"
	kv.Value = strconv.Itoa(*page)
	queryParams = append(queryParams, *kv)

	if typeID != nil {
		redisKey = redisKey + ":" + strconv.Itoa(*typeID)
		kv.Key = "type_id"
		kv.Value = strconv.Itoa(*typeID)
		queryParams = append(queryParams, *kv)
	}

	redisKey = redisKey + ":" + strconv.Itoa(*page)

	orderResult, pages, err := ordersForRegionREST(baseUrl, queryParams, redisKey)

	if err == nil {
		orderList = append(orderList, orderResult...)
	} else {
		log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID, "orderType": orderType}).Errorf("First page query for Orders has error : %v", err)
		return nil, err
	}

	returnOrders := new(model.OrderWrapper)
	returnOrders.List = orderList
	returnOrders.Xpages = &pages

	return returnOrders, nil
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

func ordersForRegionREST(url string, additionalQueryParams []configuration.Key_value, redisKey string) ([]*model.Order, int, error) {
	var orders []*model.Order
	var pages = 0
	var buffer bytes.Buffer
	responseBytes, header, err := restHelper.MakeCachingRESTCall(url, http.MethodGet, buffer, additionalQueryParams, redisKey)
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
