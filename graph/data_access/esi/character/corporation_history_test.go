package character

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

func TestSuccessfulCorporationHistory(t *testing.T) {
	jsonResponse := `[
		{
		  "corporation_id": 1000049,
		  "record_id": 37797874,
		  "start_date": "2015-04-14T18:04:00Z"
		}]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	resp, err := CorporationHistory(context.Background(), &testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName int = 1000049
	if *resp[0].CorporationID != responseName {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDCorporationHistory(t *testing.T) {
	jsonResponse := `[
		{
		  "corporation_id": 1000049,
		  "record_id": 37797874,
		  "start_date": "2015-04-14T18:04:00Z"
		}]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = nil

	_, err := CorporationHistory(context.Background(), testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCallCorporationHistory(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	_, err := CorporationHistory(context.Background(), &testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshalCorporationHistory(t *testing.T) {
	jsonResponse := `[
		{{
		  "corporation_id": 1000049,
		  "record_id": 37797874,
		  "start_date": "2015-04-14T18:04:00Z"
		}]`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	_, err := CorporationHistory(context.Background(), &testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
