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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int = new(int)
	*testID = 40176874
	var ids []*int = make([]*int, 1)
	ids[0] = testID

	resp, err := MoonDetails(context.Background(), ids)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int = new(int)
	*testID = 40176874
	var ids []*int = make([]*int, 2)
	ids[0] = testID

	_, err := MoonDetails(context.Background(), ids)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 40176874

	resp, err := MoonByID(context.Background(), &testID)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int

	_, err := MoonByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallMoonByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 40176874

	_, err := MoonByID(context.Background(), &testID)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 40176874

	_, err := MoonByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
