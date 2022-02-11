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
*              MoonDetails             *
***************************************/

func TestSuccessfulMoonDetails(t *testing.T) {
	jsonResponse := `{
		"moon_id": 40176874,
		"name": "Inaro I - Moon 1",
		"position": {
		  "x": -30422708970,
		  "y": -5774169055,
		  "z": -33009489782
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = new(int)
	*testId = 40176874
	var ids []*int = make([]*int, 1)
	ids[0] = testId

	resp, err := MoonDetails(ids, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Inaro I - Moon 1"
	if *resp[0].Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestFailNilIDMoonDetails(t *testing.T) {
	jsonResponse := `{
		"moon_id": 40176874,
		"name": "Inaro I - Moon 1",
		"position": {
		  "x": -30422708970,
		  "y": -5774169055,
		  "z": -33009489782
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = new(int)
	*testId = 40176874
	var ids []*int = make([]*int, 2)
	ids[0] = testId

	_, err := MoonDetails(ids, context.Background())
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

/***************************************
*              MoonByID                *
***************************************/

func TestSuccessfulMoonByID(t *testing.T) {
	jsonResponse := `{
		"moon_id": 40176874,
		"name": "Inaro I - Moon 1",
		"position": {
		  "x": -30422708970,
		  "y": -5774169055,
		  "z": -33009489782
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 40176874

	resp, err := MoonByID(&testId, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Inaro I - Moon 1"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDMoonByID(t *testing.T) {
	jsonResponse := `{
		"moon_id": 40176874,
		"name": "Inaro I - Moon 1",
		"position": {
		  "x": -30422708970,
		  "y": -5774169055,
		  "z": -33009489782
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = nil

	_, err := MoonByID(testId, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallMoonByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 40176874

	_, err := MoonByID(&testId, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalMoonByID(t *testing.T) {
	jsonResponse := `{{
		"moon_id": 40176874,
		"name": "Inaro I - Moon 1",
		"position": {
		  "x": -30422708970,
		  "y": -5774169055,
		  "z": -33009489782
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 40176874

	_, err := MoonByID(&testId, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	}

}
