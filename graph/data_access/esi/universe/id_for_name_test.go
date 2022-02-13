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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName = "Agent CCP Zoetrope"

	resp, err := IdForName(context.Background(), &testName, model.AGENTS)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName = "Alliance CCP Zoetrope"

	resp, err := IdForName(context.Background(), &testName, model.ALLIANCES)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName = "Character CCP Zoetrope"

	resp, err := IdForName(context.Background(), &testName, model.CHARACTERS)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName = "Constellation CCP Zoetrope"

	resp, err := IdForName(context.Background(), &testName, model.CONSTELLATIONS)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName = "Corporation CCP Zoetrope"

	resp, err := IdForName(context.Background(), &testName, model.CORPORATIONS)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName = "Faction CCP Zoetrope"

	resp, err := IdForName(context.Background(), &testName, model.FACTIONS)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName = "Inventory Type CCP Zoetrope"

	resp, err := IdForName(context.Background(), &testName, model.INVENTORY_TYPES)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName = "Region CCP Zoetrope"

	resp, err := IdForName(context.Background(), &testName, model.REGIONS)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName = "System CCP Zoetrope"

	resp, err := IdForName(context.Background(), &testName, model.SYSTEMS)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName = "System CCP Zoetrope"

	_, err := IdForName(context.Background(), &testName, "bryans")
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName *string = nil

	_, err := IdForName(context.Background(), testName, model.AGENTS)
	if err == nil {
		t.Error(helpers.NilError)
	}
}

func TestFailRESTFailureIDForName(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testName = "bryan"

	_, err := IdForName(context.Background(), &testName, model.AGENTS)
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
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testName = "Region CCP Zoetrope"

	_, err := IdForName(context.Background(), &testName, model.REGIONS)
	if err == nil {
		t.Error(helpers.NilError)
	}
}
