package universe

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

/***************************************
*          ConstellationsByIDs         *
***************************************/

func TestSuccessfulConstellationsByIDs(t *testing.T) {
	jsonResponse := `{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = new(int)
	*testId = 20000019
	var ids []*int = make([]*int, 1)
	ids[0] = testId

	resp, err := ConstellationsByIDs(ids)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Ihilakken"
	if *resp[0].Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestFailNilIDConstellationsByIDs(t *testing.T) {
	jsonResponse := `{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = new(int)
	*testId = 20000019
	var ids []*int = make([]*int, 2)
	ids[0] = testId

	_, err := ConstellationsByIDs(ids)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

/***************************************
*          ConstellationByID           *
***************************************/

func TestSuccessfulConstellationByID(t *testing.T) {
	jsonResponse := `{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 20000019

	resp, err := ConstellationByID(&testId)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Ihilakken"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDConstellationByID(t *testing.T) {
	jsonResponse := `{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = nil

	_, err := ConstellationByID(testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallConstellationByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 20000019

	_, err := ConstellationByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalConstellationByID(t *testing.T) {
	jsonResponse := `{{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 20000019

	_, err := ConstellationByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}