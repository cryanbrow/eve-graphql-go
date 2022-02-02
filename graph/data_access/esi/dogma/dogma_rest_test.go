package dogma

import (
	"bytes"
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
		MockMakeCachingRESTCall: func(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
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

type MockMakeCachingRESTCallType func(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error)

type MockRestHelper struct {
	MockMakeCachingRESTCall MockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
	return m.MockMakeCachingRESTCall(base_url, verb, body, additional_query_params, redis_query_key)
}
