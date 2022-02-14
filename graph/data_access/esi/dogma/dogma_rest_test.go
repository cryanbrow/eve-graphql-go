package dogma

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

const wrongError string = "Wrong error text: %s"

func TestSuccessfulDogmaEffectByID(t *testing.T) {
	jsonResponse := `{
		"description": "",
		"display_name": "",
		"effect_category": 0,
		"effect_id": 4513,
		"icon_id": 0,
		"modifiers": [
		  {
			"domain": "shipID",
			"func": "LocationGroupModifier",
			"modified_attribute_id": 20,
			"modifying_attribute_id": 587,
			"operator": 6
		  }
		],
		"name": "shipStasisWebStrengthBonusMF2"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		DogmaMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	resp, err := DogmaEffectByID(context.Background(), &testID)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName = "shipStasisWebStrengthBonusMF2"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDDogmaEffectByID(t *testing.T) {
	jsonResponse := `{
		"description": "",
		"display_name": "",
		"effect_category": 0,
		"effect_id": 4513,
		"icon_id": 0,
		"modifiers": [
		  {
			"domain": "shipID",
			"func": "LocationGroupModifier",
			"modified_attribute_id": 20,
			"modifying_attribute_id": 587,
			"operator": 6
		  }
		],
		"name": "shipStasisWebStrengthBonusMF2"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		DogmaMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int

	_, err := DogmaEffectByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilID {
		t.Errorf(wrongError, err.Error())
	}
}

func TestFailRestCallDogmaEffectByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		DogmaMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := DogmaEffectByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(wrongError, err.Error())
	}
}

func TestFailUnmarshalDogmaEffectByID(t *testing.T) {
	jsonResponse := `{{
		"description": "",
		"display_name": "",
		"effect_category": 0,
		"effect_id": 4513,
		"icon_id": 0,
		"modifiers": [
		  {
			"domain": "shipID",
			"func": "LocationGroupModifier",
			"modified_attribute_id": 20,
			"modifying_attribute_id": 587,
			"operator": 6
		  }
		],
		"name": "shipStasisWebStrengthBonusMF2"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		DogmaMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := DogmaEffectByID(context.Background(), &testID)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestSuccessfulDogmaAttributeByID(t *testing.T) {
	jsonResponse := `{
		"attribute_id": 1966,
		"default_value": 0,
		"description": "",
		"display_name": "Energy warfare modifier",
		"high_is_good": true,
		"name": "energyWarfareStrengthMultiplier",
		"published": true,
		"stackable": true,
		"unit_id": 104
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		DogmaMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	resp, err := DogmaAttributeByID(context.Background(), &testID)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName = "energyWarfareStrengthMultiplier"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDDogmaAttributeByID(t *testing.T) {
	jsonResponse := `{
		"attribute_id": 1966,
		"default_value": 0,
		"description": "",
		"display_name": "Energy warfare modifier",
		"high_is_good": true,
		"name": "energyWarfareStrengthMultiplier",
		"published": true,
		"stackable": true,
		"unit_id": 104
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		DogmaMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int

	_, err := DogmaAttributeByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilID {
		t.Errorf(wrongError, err.Error())
	}

}

func TestFailRestCallDogmaAttributeByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		DogmaMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := DogmaAttributeByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf(wrongError, err.Error())
	}

}

func TestFailUnmarshalDogmaAttributeByID(t *testing.T) {
	jsonResponse := `{
		"attribute_id": 1966"hi",
		"default_value": 0,
		"description": "",
		"display_name": "Energy warfare modifier",
		"high_is_good": true,
		"name": "energyWarfareStrengthMultiplier",
		"published": true,
		"stackable": true,
		"unit_id": 104
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		DogmaMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := DogmaAttributeByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

type DogmaMockMakeCachingRESTCallType func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error)

type MockRestHelper struct {
	DogmaMockMakeCachingRESTCall DogmaMockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
	return m.DogmaMockMakeCachingRESTCall(ctx, baseURL, verb, body, additionalQueryParams, redisQueryKey)
}
