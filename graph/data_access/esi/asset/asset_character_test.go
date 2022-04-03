package asset

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

func TestCorporationsSuccessfulByID(t *testing.T) {
	jsonResponse := `[
		{
		  "is_singleton": true,
		  "item_id": 1038558093296,
		  "location_flag": "Hangar",
		  "location_id": 60015037,
		  "location_type": "station",
		  "quantity": 1,
		  "type_id": 588
		},
		{
		  "is_singleton": true,
		  "item_id": 1038558163914,
		  "location_flag": "LoSlot0",
		  "location_id": 1038561617235,
		  "location_type": "item",
		  "quantity": 1,
		  "type_id": 22542
		},
		{
		  "is_blueprint_copy": true,
		  "is_singleton": true,
		  "item_id": 1038559411606,
		  "location_flag": "Hangar",
		  "location_id": 60015037,
		  "location_type": "station",
		  "quantity": 1,
		  "type_id": 1134
		}
		]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		AssetMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 2119771486

	resp, err := AssetsByCharacterID(context.Background(), &testID)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}

	var firstAsset = *resp[0]

	if *firstAsset.ItemID != 1038558093296 {
		t.Errorf("Response was not as expected")
	}

}

func TestCorporationsFailNilIDByID(t *testing.T) {
	jsonResponse := `[
		{
		  "is_singleton": true,
		  "item_id": 1038558093296,
		  "location_flag": "Hangar",
		  "location_id": 60015037,
		  "location_type": "station",
		  "quantity": 1,
		  "type_id": 588
		},
		{
		  "is_singleton": true,
		  "item_id": 1038558163914,
		  "location_flag": "LoSlot0",
		  "location_id": 1038561617235,
		  "location_type": "item",
		  "quantity": 1,
		  "type_id": 22542
		},
		{
		  "is_blueprint_copy": true,
		  "is_singleton": true,
		  "item_id": 1038559411606,
		  "location_flag": "Hangar",
		  "location_id": 60015037,
		  "location_type": "station",
		  "quantity": 1,
		  "type_id": 1134
		}
		]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		AssetMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int

	_, err := AssetsByCharacterID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilID {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestCorporationsFailRestCallByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		AssetMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := AssetsByCharacterID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestCorporationsFailUnmarshalByID(t *testing.T) {
	jsonResponse := `{[
		207315351
	  ]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		AssetMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := AssetsByCharacterID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

type CorporationMockMakeCachingRESTCallType func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error)

type MockCorporationRestHelper struct {
	CorporationMockMakeCachingRESTCall CorporationMockMakeCachingRESTCallType
}

func (m *MockCorporationRestHelper) MakeCachingRESTCall(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
	return m.CorporationMockMakeCachingRESTCall(ctx, baseURL, verb, body, additionalQueryParams, redisQueryKey)
}
