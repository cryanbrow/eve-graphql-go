package corporation

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

func TestSuccessfulCorporationByID(t *testing.T) {
	jsonResponse := `{
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

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CorporationMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId = 1

	resp, err := CorporationByID(context.Background(), &testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Tribal Liberation Force"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDCorporationByID(t *testing.T) {
	jsonResponse := `{
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

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CorporationMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int

	_, err := CorporationByID(context.Background(), testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCallCorporationByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		CorporationMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId = 1

	_, err := CorporationByID(context.Background(), &testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshalCorporationByID(t *testing.T) {
	jsonResponse := `{{
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

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CorporationMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId = 1

	_, err := CorporationByID(context.Background(), &testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

type CorporationMockMakeCachingRESTCallType func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error)

type MockRestHelper struct {
	CorporationMockMakeCachingRESTCall CorporationMockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
	return m.CorporationMockMakeCachingRESTCall(ctx, baseURL, verb, body, additionalQueryParams, redisQueryKey)
}
