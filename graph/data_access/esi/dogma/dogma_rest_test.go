package dogma

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

func TestSuccessful_DogmaEffectByID(t *testing.T) {
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

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	resp, err := DogmaEffectByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "shipStasisWebStrengthBonusMF2"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilID_DogmaEffectByID(t *testing.T) {
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

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id *int = nil

	_, err := DogmaEffectByID(test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "nil id" {
		t.Errorf("Wrong error text: %s", err.Error())
	}
}

func TestFailRestCall_DogmaEffectByID(t *testing.T) {
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	_, err := DogmaEffectByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}
}

func TestFailUnmarshal_DogmaEffectByID(t *testing.T) {
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

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	_, err := DogmaEffectByID(&test_id)
	if err == nil {
		t.Errorf("Error is nil")
	}
}

func TestSuccessful_DogmaAttributeByID(t *testing.T) {
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

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	resp, err := DogmaAttributeByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "energyWarfareStrengthMultiplier"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilID_DogmaAttributeByID(t *testing.T) {
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

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id *int = nil

	_, err := DogmaAttributeByID(test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "nil id" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCall_DogmaAttributeByID(t *testing.T) {
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	_, err := DogmaAttributeByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshal_DogmaAttributeByID(t *testing.T) {
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

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	_, err := DogmaAttributeByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	}

}

type MockMakeCachingRESTCallType func(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error)

type MockRestHelper struct {
	MockMakeCachingRESTCall MockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
	return m.MockMakeCachingRESTCall(baseUrl, verb, body, additional_query_params, redis_query_key)
}
