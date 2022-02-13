package market

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/generated/model"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

/***************************************
*             OrdersForRegion              *
***************************************/

func TestSuccessfulOrdersForRegion(t *testing.T) {
	jsonResponse := `[
		{
		  "duration": 90,
		  "is_buy_order": true,
		  "issued": "2022-01-30T18:47:21Z",
		  "location_id": 1035026346682,
		  "min_volume": 1,
		  "order_id": 6161383608,
		  "price": 2280000,
		  "range": "40",
		  "system_id": 30000732,
		  "type_id": 44992,
		  "volume_remain": 155,
		  "volume_total": 500
		}]`

	b := []byte(jsonResponse)

	header := http.Header{
		"x-pages": []string{"1"},
	}

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, header, nil
		},
	}
	restHelper = mockRestHelper

	var regionId = 10000008
	var typeId = 44992
	var page = 1
	var orderType model.Ordertype = model.OrdertypeAll

	resp, err := OrdersForRegion(context.Background(), &regionId, &orderType, &typeId, &page)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName int = 30000732
	if *resp.List[0].SystemID != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilTypeIDOrdersForRegion(t *testing.T) {
	jsonResponse := `[
		{
		  "duration": 90,
		  "is_buy_order": true,
		  "issued": "2022-01-30T18:47:21Z",
		  "location_id": 1035026346682,
		  "min_volume": 1,
		  "order_id": 6161383608,
		  "price": 2280000,
		  "range": "40",
		  "system_id": 30000732,
		  "type_id": 44992,
		  "volume_remain": 155,
		  "volume_total": 500
		}]`

	b := []byte(jsonResponse)

	header := http.Header{
		"x-pages": []string{"5"},
	}

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, header, nil
		},
	}
	restHelper = mockRestHelper

	var regionId = 10000008
	var typeId *int
	var page = 1
	var orderType model.Ordertype = model.OrdertypeAll

	resp, err := OrdersForRegion(context.Background(), &regionId, &orderType, typeId, &page)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName int = 30000732
	if *resp.List[0].SystemID != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailRestCallOrdersForRegion(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var regionId = 10000008
	var typeId = 44992
	var page = 1
	var orderType model.Ordertype = model.OrdertypeAll

	_, err := OrdersForRegion(context.Background(), &regionId, &orderType, &typeId, &page)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalOrdersForRegion(t *testing.T) {
	jsonResponse := `[{
		{
		  "duration": 90,
		  "is_buy_order": true,
		  "issued": "2022-01-30T18:47:21Z",
		  "location_id": 1035026346682,
		  "min_volume": 1,
		  "order_id": 6161383608,
		  "price": 2280000,
		  "range": "40",
		  "system_id": 30000732,
		  "type_id": 44992,
		  "volume_remain": 155,
		  "volume_total": 500
		}]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var regionId = 10000008
	var typeId = 44992
	var page = 1
	var orderType model.Ordertype = model.OrdertypeAll

	_, err := OrdersForRegion(context.Background(), &regionId, &orderType, &typeId, &page)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
