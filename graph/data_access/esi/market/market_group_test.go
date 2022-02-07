package market

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

/***************************************
*             MarketGroupByID              *
***************************************/

func TestSuccessfulMarketGroupByID(t *testing.T) {
	jsonResponse := `{
		"description": "Blueprints are data items used in industry for manufacturing, research and invention jobs",
		"market_group_id": 2,
		"name": "Blueprints & Reactions",
		"types": []
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 2

	resp, err := MarketGroupByID(&testId)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Blueprints & Reactions"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDMarketGroupByID(t *testing.T) {
	jsonResponse := `{
		"description": "Blueprints are data items used in industry for manufacturing, research and invention jobs",
		"market_group_id": 2,
		"name": "Blueprints & Reactions",
		"types": []
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = nil

	_, err := MarketGroupByID(testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallMarketGroupByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 2

	_, err := MarketGroupByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalMarketGroupByID(t *testing.T) {
	jsonResponse := `{{
		"description": "Blueprints are data items used in industry for manufacturing, research and invention jobs",
		"market_group_id": 2,
		"name": "Blueprints & Reactions",
		"types": []
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 2

	_, err := MarketGroupByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

type MockMakeCachingRESTCallType func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error)

type MockRestHelper struct {
	MockMakeCachingRESTCall MockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
	return m.MockMakeCachingRESTCall(baseUrl, verb, body, additionalQueryParams, redisQueryKey)
}
