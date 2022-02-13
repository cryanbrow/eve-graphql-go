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
*              StargateDetails             *
***************************************/

func TestSuccessfulStargateDetails(t *testing.T) {
	jsonResponse := `{
		"destination": {
		  "stargate_id": 50003584,
		  "system_id": 30002789
		},
		"name": "Stargate (Kaaputenen)",
		"position": {
		  "x": -1674740736000,
		  "y": -317977681920,
		  "z": -2212440760320
		},
		"stargate_id": 50003085,
		"system_id": 30002788,
		"type_id": 16
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = new(int)
	*testID = 50003085
	var ids = make([]*int, 1)
	ids[0] = testID

	resp, err := StargateDetails(context.Background(), ids)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Stargate (Kaaputenen)"
	if *resp[0].Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestFailNilIDStargateDetails(t *testing.T) {
	jsonResponse := `{
		"destination": {
		  "stargate_id": 50003584,
		  "system_id": 30002789
		},
		"name": "Stargate (Kaaputenen)",
		"position": {
		  "x": -1674740736000,
		  "y": -317977681920,
		  "z": -2212440760320
		},
		"stargate_id": 50003085,
		"system_id": 30002788,
		"type_id": 16
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = new(int)
	*testID = 50003085
	var ids = make([]*int, 2)
	ids[0] = testID

	_, err := StargateDetails(context.Background(), ids)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

/***************************************
*              StargateByID                *
***************************************/

func TestSuccessfulStargateByID(t *testing.T) {
	jsonResponse := `{
		"destination": {
		  "stargate_id": 50003584,
		  "system_id": 30002789
		},
		"name": "Stargate (Kaaputenen)",
		"position": {
		  "x": -1674740736000,
		  "y": -317977681920,
		  "z": -2212440760320
		},
		"stargate_id": 50003085,
		"system_id": 30002788,
		"type_id": 16
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 50003085

	resp, err := StargateByID(context.Background(), &testID)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Stargate (Kaaputenen)"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}

}

func TestFailNilIDStargateByID(t *testing.T) {
	jsonResponse := `{
		"destination": {
		  "stargate_id": 50003584,
		  "system_id": 30002789
		},
		"name": "Stargate (Kaaputenen)",
		"position": {
		  "x": -1674740736000,
		  "y": -317977681920,
		  "z": -2212440760320
		},
		"stargate_id": 50003085,
		"system_id": 30002788,
		"type_id": 16
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int

	_, err := StargateByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilID {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailRestCallStargateByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 50003085

	_, err := StargateByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(helpers.WrongErrorText, err.Error())
	}

}

func TestFailUnmarshalStargateByID(t *testing.T) {
	jsonResponse := `{{
		"destination": {
		  "stargate_id": 50003584,
		  "system_id": 30002789
		},
		"name": "Stargate (Kaaputenen)",
		"position": {
		  "x": -1674740736000,
		  "y": -317977681920,
		  "z": -2212440760320
		},
		"stargate_id": 50003085,
		"system_id": 30002788,
		"type_id": 16
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 50003085

	_, err := StargateByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
