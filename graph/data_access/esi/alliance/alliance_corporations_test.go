package alliance

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/data_access/esi/corporation"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

func TestCorporationsSuccessfulByID(t *testing.T) {
	jsonResponse := `[
		207315351
	  ]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		AllianceMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	jsonCorporationResponse := `{
		"ceo_id": 3018996,
		"creator_id": 1,
		"description": "The Minmatar heart sings for freedom and the Minmatar soul strives for open skies, but the Minmatar heart has been withering away in captivity. It is now up to you, capsuleer. You hold the power to free our people. You are the heroes of your generation. Join us in the struggle for freedom. Death to Amarr; long live the Minmatar Nation.",
		"faction_id": 500002,
		"home_station_id": 60015096,
		"member_count": 22862,
		"name": "Tribal Liberation Force",
		"shares": 0,
		"tax_rate": 0,
		"ticker": "TLIB",
		"url": ""
	  }`

	b2 := []byte(jsonCorporationResponse)

	mockCorporationRestHelper := &MockRestHelper{
		AllianceMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b2, nil, nil
		},
	}
	corporation.RESTHelper = mockCorporationRestHelper

	var testID = 99000068

	resp, err := CorporationsByID(context.Background(), &testID)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}

	var localCorp = *resp[0]
	var ceoID = *localCorp.CeoID

	if ceoID != 3018996 {
		t.Errorf("Response was not as expected")
	}

}

func TestCorporationsFailNilIDByID(t *testing.T) {
	jsonResponse := `[
		207315351
	  ]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		AllianceMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int

	_, err := CorporationsByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilID {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestCorporationsFailRestCallByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		AllianceMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := CorporationsByID(context.Background(), &testID)
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
		AllianceMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := CorporationsByID(context.Background(), &testID)
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
