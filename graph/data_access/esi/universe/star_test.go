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
*               StarByID               *
***************************************/

func TestSuccessfulStarByID(t *testing.T) {
	jsonResponse := `{
		"age": 1272593654,
		"luminosity": 0.4643999934196472,
		"name": "Inaro - Star",
		"radius": 483100000,
		"solar_system_id": 30002788,
		"spectral_class": "G5 V",
		"temperature": 5449,
		"type_id": 3797
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 40176872

	resp, err := StarByID(&testId, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Inaro - Star"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDStarByID(t *testing.T) {
	var testId *int = nil

	_, err := StarByID(testId, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallStarByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 40176872

	_, err := StarByID(&testId, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalStarByID(t *testing.T) {
	jsonResponse := `{{
		"age": 1272593654,
		"luminosity": 0.4643999934196472,
		"name": "Inaro - Star",
		"radius": 483100000,
		"solar_system_id": 30002788,
		"spectral_class": "G5 V",
		"temperature": 5449,
		"type_id": 3797
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 40176872

	_, err := StarByID(&testId, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	}

}
