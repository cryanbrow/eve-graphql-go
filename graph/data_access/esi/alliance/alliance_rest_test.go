package alliance

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

func TestSuccessfulAllianceByID(t *testing.T) {
	jsonResponse := `{
		"creator_corporation_id": 98007669,
		"creator_id": 1973270502,
		"date_founded": "2010-12-13T02:49:00Z",
		"executor_corporation_id": 296119337,
		"name": "Fleet Coordination Coalition",
		"ticker": "FC0RD"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		AllianceMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	resp, err := AllianceByID(context.Background(), &testID)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName = "Fleet Coordination Coalition"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDAllianceByID(t *testing.T) {
	jsonResponse := `{
		"creator_corporation_id": 98007669,
		"creator_id": 1973270502,
		"date_founded": "2010-12-13T02:49:00Z",
		"executor_corporation_id": 296119337,
		"name": "Fleet Coordination Coalition",
		"ticker": "FC0RD"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		AllianceMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int

	_, err := AllianceByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilID {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCallAllianceByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		AllianceMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := AllianceByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshalAllianceByID(t *testing.T) {
	jsonResponse := `{{
		"creator_corporation_id": 98007669,
		"creator_id": 1973270502,
		"date_founded": "2010-12-13T02:49:00Z",
		"executor_corporation_id": 296119337,
		"name": "Fleet Coordination Coalition",
		"ticker": "FC0RD"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		AllianceMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := AllianceByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

type AllianceMockMakeCachingRESTCallType func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error)

type MockRestHelper struct {
	AllianceMockMakeCachingRESTCall AllianceMockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
	return m.AllianceMockMakeCachingRESTCall(ctx, baseURL, verb, body, additionalQueryParams, redisQueryKey)
}
