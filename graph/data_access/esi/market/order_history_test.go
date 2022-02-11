package market

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

/***************************************
*             OrderHistory              *
***************************************/

func TestSuccessfulOrderHistory(t *testing.T) {
	jsonResponse := `[
		{
		  "average": 700000,
		  "date": "2021-05-16",
		  "highest": 700000,
		  "lowest": 700000,
		  "order_count": 1,
		  "volume": 1
		}]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var regionId int = 10000008
	var typeId int = 602

	resp, err := OrderHistory(&regionId, &typeId, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseVolume int = 1
	if *resp[0].Volume != responseVolume {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilRegionIDOrderHistory(t *testing.T) {
	jsonResponse := `[
		{
		  "average": 700000,
		  "date": "2021-05-16",
		  "highest": 700000,
		  "lowest": 700000,
		  "order_count": 1,
		  "volume": 1
		}]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var regionId *int = nil
	var typeId int = 602

	_, err := OrderHistory(regionId, &typeId, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailNilTypeIDOrderHistory(t *testing.T) {
	jsonResponse := `[
		{
		  "average": 700000,
		  "date": "2021-05-16",
		  "highest": 700000,
		  "lowest": 700000,
		  "order_count": 1,
		  "volume": 1
		}]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var regionId int = 10000008
	var typeId *int = nil

	_, err := OrderHistory(&regionId, typeId, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallOrderHistory(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var regionId int = 2
	var typeId int = 2

	_, err := OrderHistory(&regionId, &typeId, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalOrderHistory(t *testing.T) {
	jsonResponse := `[{
		{
		  "average": 700000,
		  "date": "2021-05-16",
		  "highest": 700000,
		  "lowest": 700000,
		  "order_count": 1,
		  "volume": 1
		}]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var regionId int = 10000008
	var typeId int = 602

	_, err := OrderHistory(&regionId, &typeId, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	}

}
