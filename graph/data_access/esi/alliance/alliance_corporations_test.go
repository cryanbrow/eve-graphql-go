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

	var testID = 99000068

	resp, err := CorporationsByID(context.Background(), &testID)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName = 207315351

	if (*resp)[0] != responseName {
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
