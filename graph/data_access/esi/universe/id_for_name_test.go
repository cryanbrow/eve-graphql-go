package universe

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
	"github.com/cryanbrow/eve-graphql-go/graph/model"
)

func TestSuccessfulAgentIDForName(t *testing.T) {
	jsonResponse := `{
		"agents": [
		  {
			"id": 2112625428,
			"name": "Agent CCP Zoetrope"
		  }
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName string = "Agent CCP Zoetrope"

	resp, err := IdForName(&testName, model.AGENTS, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseInt int = 2112625428
	if resp != responseInt {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulAllianceIDForName(t *testing.T) {
	jsonResponse := `{
		"alliances": [
		  {
			"id": 2112625428,
			"name": "Alliance CCP Zoetrope"
		  }
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName string = "Alliance CCP Zoetrope"

	resp, err := IdForName(&testName, model.ALLIANCES, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseInt int = 2112625428
	if resp != responseInt {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulCharacterIDForName(t *testing.T) {
	jsonResponse := `{
		"characters": [
		  {
			"id": 2112625428,
			"name": "Character CCP Zoetrope"
		  }
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName string = "Character CCP Zoetrope"

	resp, err := IdForName(&testName, model.CHARACTERS, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseInt int = 2112625428
	if resp != responseInt {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulConstellationIDForName(t *testing.T) {
	jsonResponse := `{
		"constellations": [
		  {
			"id": 2112625428,
			"name": "Constellation CCP Zoetrope"
		  }
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName string = "Constellation CCP Zoetrope"

	resp, err := IdForName(&testName, model.CONSTELLATIONS, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseInt int = 2112625428
	if resp != responseInt {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulCorporationIDForName(t *testing.T) {
	jsonResponse := `{
		"corporations": [
		  {
			"id": 2112625428,
			"name": "Corporation CCP Zoetrope"
		  }
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName string = "Corporation CCP Zoetrope"

	resp, err := IdForName(&testName, model.CORPORATIONS, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseInt int = 2112625428
	if resp != responseInt {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulFactionIDForName(t *testing.T) {
	jsonResponse := `{
		"factions": [
		  {
			"id": 2112625428,
			"name": "Faction CCP Zoetrope"
		  }
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName string = "Faction CCP Zoetrope"

	resp, err := IdForName(&testName, model.FACTIONS, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseInt int = 2112625428
	if resp != responseInt {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulInventoryTypeIDForName(t *testing.T) {
	jsonResponse := `{
		"inventory_types": [
		  {
			"id": 2112625428,
			"name": "Inventory Type CCP Zoetrope"
		  }
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName string = "Inventory Type CCP Zoetrope"

	resp, err := IdForName(&testName, model.INVENTORY_TYPES, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseInt int = 2112625428
	if resp != responseInt {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulRegionIDForName(t *testing.T) {
	jsonResponse := `{
		"regions": [
		  {
			"id": 2112625428,
			"name": "Region CCP Zoetrope"
		  }
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName string = "Region CCP Zoetrope"

	resp, err := IdForName(&testName, model.REGIONS, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseInt int = 2112625428
	if resp != responseInt {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulSystemIDForName(t *testing.T) {
	jsonResponse := `{
		"systems": [
		  {
			"id": 2112625428,
			"name": "System CCP Zoetrope"
		  }
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName string = "System CCP Zoetrope"

	resp, err := IdForName(&testName, model.SYSTEMS, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseInt int = 2112625428
	if resp != responseInt {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestFailNilNameTypeIDForName(t *testing.T) {
	jsonResponse := `{
		"bryans": [
		  {
			"id": 2112625428,
			"name": "System CCP Zoetrope"
		  }
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName string = "System CCP Zoetrope"

	_, err := IdForName(&testName, "bryans", context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	}
}

func TestFailNilNameIDForName(t *testing.T) {
	jsonResponse := `{
		"agents": [
		  {
			"id": 2112625428,
			"name": "System CCP Zoetrope"
		  }
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName *string = nil

	_, err := IdForName(testName, model.AGENTS, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	}
}

func TestFailRESTFailureIDForName(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testName string = "bryan"

	_, err := IdForName(&testName, model.AGENTS, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	}
}

func TestFailureUnmarshalIDForName(t *testing.T) {
	jsonResponse := `{{
		"regions": [
		  {
			"id": 2112625428,
			"name": "Region CCP Zoetrope"
		  }
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName string = "Region CCP Zoetrope"

	_, err := IdForName(&testName, model.REGIONS, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	}
}
