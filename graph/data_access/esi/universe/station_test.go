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
*            StationByID               *
***************************************/

func TestSuccessfulStationByID(t *testing.T) {
	jsonResponse := `{
		"max_dockable_ship_volume": 50000000,
		"name": "Ebolfer V - Tribal Liberation Force Testing Facilities",
		"office_rental_cost": 10000,
		"owner": 1000182,
		"position": {
		  "x": -30282670080,
		  "y": 5940387840,
		  "z": -1234189639680
		},
		"race_id": 2,
		"reprocessing_efficiency": 0.5,
		"reprocessing_stations_take": 0.025,
		"services": [
		  "bounty-missions",
		  "courier-missions",
		  "reprocessing-plant",
		  "market",
		  "repair-facilities",
		  "factory",
		  "labratory",
		  "fitting",
		  "news",
		  "insurance",
		  "docking",
		  "office-rental",
		  "loyalty-point-store",
		  "navy-offices",
		  "security-offices"
		],
		"station_id": 60015096,
		"system_id": 30002094,
		"type_id": 2500
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 60015096

	resp, err := StationByID(&testId)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Ebolfer V - Tribal Liberation Force Testing Facilities"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDStationByID(t *testing.T) {
	var testId *int = nil

	_, err := StationByID(testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallStationByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 60015096

	_, err := StationByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalStationByID(t *testing.T) {
	jsonResponse := `{{
		"max_dockable_ship_volume": 50000000,
		"name": "Ebolfer V - Tribal Liberation Force Testing Facilities",
		"office_rental_cost": 10000,
		"owner": 1000182,
		"position": {
		  "x": -30282670080,
		  "y": 5940387840,
		  "z": -1234189639680
		},
		"race_id": 2,
		"reprocessing_efficiency": 0.5,
		"reprocessing_stations_take": 0.025,
		"services": [
		  "bounty-missions",
		  "courier-missions",
		  "reprocessing-plant",
		  "market",
		  "repair-facilities",
		  "factory",
		  "labratory",
		  "fitting",
		  "news",
		  "insurance",
		  "docking",
		  "office-rental",
		  "loyalty-point-store",
		  "navy-offices",
		  "security-offices"
		],
		"station_id": 60015096,
		"system_id": 30002094,
		"type_id": 2500
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 60015096

	_, err := StationByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

func TestSuccessfulStationsByIDs(t *testing.T) {
	jsonResponse := `{
		"max_dockable_ship_volume": 50000000,
		"name": "Ebolfer V - Tribal Liberation Force Testing Facilities",
		"office_rental_cost": 10000,
		"owner": 1000182,
		"position": {
		  "x": -30282670080,
		  "y": 5940387840,
		  "z": -1234189639680
		},
		"race_id": 2,
		"reprocessing_efficiency": 0.5,
		"reprocessing_stations_take": 0.025,
		"services": [
		  "bounty-missions",
		  "courier-missions",
		  "reprocessing-plant",
		  "market",
		  "repair-facilities",
		  "factory",
		  "labratory",
		  "fitting",
		  "news",
		  "insurance",
		  "docking",
		  "office-rental",
		  "loyalty-point-store",
		  "navy-offices",
		  "security-offices"
		],
		"station_id": 60015096,
		"system_id": 30002094,
		"type_id": 2500
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId1 int = 60015096
	var testId2 int = 60015096
	var ids []*int = make([]*int, 2)
	ids[0] = &testId1
	ids[1] = &testId2

	resp, err := StationsByIDs(ids)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Ebolfer V - Tribal Liberation Force Testing Facilities"
	if *resp[0].Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestFailureStationsByIDs(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId1 int = 60015096
	var testId2 int = 60015096
	var ids []*int = make([]*int, 2)
	ids[0] = &testId1
	ids[1] = &testId2

	_, err := StationsByIDs(ids)
	if err == nil {
		t.Error(helpers.NilError)
	}
}
