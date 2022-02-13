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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 60015096

	resp, err := StationByID(context.Background(), &testID)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Ebolfer V - Tribal Liberation Force Testing Facilities"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDStationByID(t *testing.T) {
	var testID *int

	_, err := StationByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilID {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallStationByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 60015096

	_, err := StationByID(context.Background(), &testID)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 60015096

	_, err := StationByID(context.Background(), &testID)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID1 int = 60015096
	var testID2 int = 60015096
	var ids []*int = make([]*int, 2)
	ids[0] = &testID1
	ids[1] = &testID2

	resp, err := StationsByIDs(context.Background(), ids)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Ebolfer V - Tribal Liberation Force Testing Facilities"
	if *resp[0].Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestFailureStationsByIDs(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID1 int = 60015096
	var testID2 int = 60015096
	var ids []*int = make([]*int, 2)
	ids[0] = &testID1
	ids[1] = &testID2

	_, err := StationsByIDs(context.Background(), ids)
	if err == nil {
		t.Error(helpers.NilError)
	}
}
