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
		MockMakeCachingRESTCall: func(ctx context.Context, string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId = 2

	resp, err := MarketGroupByID(context.Background(), &testId)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int

	_, err := MarketGroupByID(context.Background(), testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallMarketGroupByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId = 2

	_, err := MarketGroupByID(context.Background(), &testId)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId = 2

	_, err := MarketGroupByID(context.Background(), &testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
