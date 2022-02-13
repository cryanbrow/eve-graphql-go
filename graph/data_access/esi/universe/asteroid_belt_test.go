package universe

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
*          AsteroidBeltDetails         *
***************************************/

func TestSuccessfulAsteroidBeltDetails(t *testing.T) {
	jsonResponse := `{
		"name": "Inaro IX - Asteroid Belt 1",
		"position": {
		  "x": 809389301760,
		  "y": 151954759680,
		  "z": -221539000320
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = new(int)
	*testID = 1
	var ids = make([]*int, 1)
	ids[0] = testID

	resp, err := AsteroidBeltDetails(context.Background(), ids)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Inaro IX - Asteroid Belt 1"
	if *resp[0].Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestFailNilIDAsteroidBeltDetails(t *testing.T) {
	jsonResponse := `{
		"name": "Inaro IX - Asteroid Belt 1",
		"position": {
		  "x": 809389301760,
		  "y": 151954759680,
		  "z": -221539000320
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = new(int)
	*testID = 1
	var ids = make([]*int, 2)
	ids[0] = testID

	_, err := AsteroidBeltDetails(context.Background(), ids)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

/***************************************
*           AsteroidBeltByID           *
***************************************/

func TestSuccessfulAsteroidBeltByID(t *testing.T) {
	jsonResponse := `{
		"name": "Inaro IX - Asteroid Belt 1",
		"position": {
		  "x": 809389301760,
		  "y": 151954759680,
		  "z": -221539000320
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	resp, err := AsteroidBeltByID(context.Background(), &testID)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Inaro IX - Asteroid Belt 1"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDAsteroidBeltByID(t *testing.T) {
	jsonResponse := `{
		"name": "Inaro IX - Asteroid Belt 1",
		"position": {
		  "x": 809389301760,
		  "y": 151954759680,
		  "z": -221539000320
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int

	_, err := AsteroidBeltByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilID {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallAsteroidBeltByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := AsteroidBeltByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalAsteroidBeltByID(t *testing.T) {
	jsonResponse := `{{
		"name": "Inaro IX - Asteroid Belt 1",
		"position": {
		  "x": 809389301760,
		  "y": 151954759680,
		  "z": -221539000320
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := AsteroidBeltByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
