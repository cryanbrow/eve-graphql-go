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

func TestSuccessfulIconByID(t *testing.T) {
	jsonResponse := `{
		"px128x128": "https://images.evetech.net/Alliance/99000036_128.png",
		"px64x64": "https://images.evetech.net/Alliance/99000036_64.png"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		AllianceMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	resp, err := IconByID(context.Background(), &testID)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	if resp == nil {
		t.Errorf("Response was nil")
	}

}

func TestFailNilIDIconByID(t *testing.T) {
	jsonResponse := `{
		"px128x128": "https://images.evetech.net/Alliance/99000036_128.png",
		"px64x64": "https://images.evetech.net/Alliance/99000036_64.png"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		AllianceMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int

	_, err := IconByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilID {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCallIconByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		AllianceMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := IconByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshalIconByID(t *testing.T) {
	jsonResponse := `{{
		"px128x128": "https://images.evetech.net/Alliance/99000036_128.png",
		"px64x64": "https://images.evetech.net/Alliance/99000036_64.png"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		AllianceMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := IconByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
