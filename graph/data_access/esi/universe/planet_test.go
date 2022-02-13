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
*              PlanetByID              *
***************************************/

func TestSuccessfulPlanetByID(t *testing.T) {
	jsonResponse := `{
		"name": "Inaro III",
		"planet_id": 40176876,
		"position": {
		  "x": -63378430212,
		  "y": -12029098752,
		  "z": 23948390844
		},
		"system_id": 30002788,
		"type_id": 2015
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 40176876

	resp, err := PlanetByID(context.Background(), &testId)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Inaro III"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDPlanetByID(t *testing.T) {
	var testId *int = nil

	_, err := PlanetByID(context.Background(), testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallPlanetByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 40176876

	_, err := PlanetByID(context.Background(), &testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalPlanetByID(t *testing.T) {
	jsonResponse := `{{
		"name": "Inaro III",
		"planet_id": 40176876,
		"position": {
		  "x": -63378430212,
		  "y": -12029098752,
		  "z": 23948390844
		},
		"system_id": 30002788,
		"type_id": 2015
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 40176876

	_, err := PlanetByID(context.Background(), &testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
