package market

import (
	"bytes"
	"context"
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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func OrdersForRegion(ctx context.Context, regionID *int, orderType *model.Ordertype, typeID *int, page *int) (*model.OrderWrapper, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "OrdersForRegion")
	defer span.End()
	log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID, "orderType": orderType}).Info("OrdersForRegion Called")
	orderList := make([]*model.Order, 0)
	baseURL := fmt.Sprintf("%s/markets/%s/orders/", configuration.AppConfig.Esi.Default.URL, strconv.Itoa(*regionID))

	redisKey := "OrdersForRegion:" + strconv.Itoa(*regionID) + ":" + orderType.String()

	queryParams := make([]configuration.KeyValue, 2)
	kv := new(configuration.KeyValue)
	kv.Key = "page"
	kv.Value = strconv.Itoa(*page)
	queryParams = append(queryParams, *kv)

	kv2 := new(configuration.KeyValue)
	kv2.Key = "order_type"
	kv2.Value = orderType.String()
	queryParams = append(queryParams, *kv2)

	if typeID != nil {
		redisKey = redisKey + ":" + strconv.Itoa(*typeID)
		kv.Key = "type_id"
		kv.Value = strconv.Itoa(*typeID)
		queryParams = append(queryParams, *kv)
	}

	redisKey = redisKey + ":" + strconv.Itoa(*page)
	log.Debugf("Here: %s", redisKey)

	orderResult, pages, err := ordersForRegionREST(newCtx, baseURL, queryParams, redisKey)

	if err == nil {
		orderList = append(orderList, orderResult...)
	} else {
		log.WithFields(log.Fields{"regionID": regionID, "typeID": typeID, "orderType": orderType}).Errorf("First page query for Orders has error : %v", err)
		return nil, err
	}

	returnOrders := new(model.OrderWrapper)
	returnOrders.List = orderList
	returnOrders.Xpages = &pages
	span.SetAttributes(attribute.Int("regionID", *regionID), attribute.Int("page", *page))
	return returnOrders, nil
}

func OrdersForRegionByName(ctx context.Context, region *string, orderType *model.Ordertype, typeName *string, page *int) (*model.OrderWrapper, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "OrdersForRegionByName")
	defer span.End()
	regionID, err := universe.IDForName(newCtx, region, local_model.Regions)
	if err != nil {
		return nil, errors.New("unknown name for region")
	}
	typeID, err := universe.IDForName(newCtx, typeName, local_model.InventoryTypes)
	if err != nil {
		return nil, errors.New("unknown name for typeName")
	}
	orders, err := OrdersForRegion(newCtx, &regionID, orderType, &typeID, page)
	if err != nil {
		return nil, err
	}

	span.SetAttributes(attribute.String("regionName", *region), attribute.String("typeName", *typeName), attribute.Int("page", *page))
	return orders, nil
}

func ordersForRegionREST(ctx context.Context, url string, additionalQueryParams []configuration.KeyValue, redisKey string) ([]*model.Order, int, error) {
	newCtx, span := otel.Tracer(tracerName).Start(ctx, "ordersForRegionREST")
	defer span.End()
	var orders []*model.Order
	var pages int
	var buffer bytes.Buffer
	responseBytes, header, err := restHelper.MakeCachingRESTCall(newCtx, url, http.MethodGet, buffer, additionalQueryParams, redisKey)
	if err != nil {
		return orders, 0, err
	}

	pages, _ = strconv.Atoi(header.Get("x-pages"))

	if err := json.Unmarshal(responseBytes, &orders); err != nil {
		log.WithFields(log.Fields{"url": url}).Errorf("Could not unmarshal reponseBytes. : %v", err.Error())
		return orders, 0, err
	}

	span.SetAttributes(attribute.String("baseURL", url), attribute.String("redisKey", redisKey))
	return orders, pages, nil
}
