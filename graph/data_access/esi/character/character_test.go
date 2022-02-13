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

func TestSuccessfulCharacterByID(t *testing.T) {
	jsonResponse := `{
		"birthday": "2008-04-01T10:24:00Z",
		"bloodline_id": 3,
		"corporation_id": 1000182,
		"description": "",
		"gender": "male",
		"name": "Gar Orga",
		"race_id": 2,
		"security_status": 0
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId = 1

	resp, err := CharacterByID(context.Background(), &testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Gar Orga"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDCharacterByID(t *testing.T) {
	jsonResponse := `{
		"birthday": "2008-04-01T10:24:00Z",
		"bloodline_id": 3,
		"corporation_id": 1000182,
		"description": "",
		"gender": "male",
		"name": "Gar Orga",
		"race_id": 2,
		"security_status": 0
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int

	_, err := CharacterByID(context.Background(), testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCallCharacterByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId = 1

	_, err := CharacterByID(context.Background(), &testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshalCharacterByID(t *testing.T) {
	jsonResponse := `{{
		"birthday": "2008-04-01T10:24:00Z",
		"bloodline_id": 3,
		"corporation_id": 1000182,
		"description": "",
		"gender": "male",
		"name": "Gar Orga",
		"race_id": 2,
		"security_status": 0
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId = 1

	_, err := CharacterByID(context.Background(), &testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
