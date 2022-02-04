package universe

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

var ancestriesJsonResponse string = `[
	{
	  "bloodline_id": 7,
	  "description": "The Gallente prize political activism more so than other Empires. Many devote their efforts towards one or more causes that suit their ambitions. Activists understand that things will never change for the better unless someone has the courage to fight the good fight.",
	  "icon_id": 1653,
	  "id": 13,
	  "name": "Activists",
	  "short_description": "Making the universe a better place, one fight at a time."
	}
  ]`

func TestSuccessfulInCache_AncestryByID(t *testing.T) {
	jsonResponse := `{
		"bloodline_id": 7,
		"description": "The Gallente prize political activism more so than other Empires. Many devote their efforts towards one or more causes that suit their ambitions. Activists understand that things will never change for the better unless someone has the courage to fight the good fight.",
		"icon_id": 1653,
		"id": 13,
		"name": "Activists",
		"short_description": "Making the universe a better place, one fight at a time."
	  }`
	b := []byte(jsonResponse)

	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return true, b
		},
	}
	Redis_client = mock_redis_client

	var test_id int = 1
	resp, err := AncestryByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Activists"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}
}

func TestSuccessfulNotInCache_AncestryByID(t *testing.T) {
	b := []byte(ancestriesJsonResponse)
	mock_redis_client := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}

	Redis_client = mock_redis_client
	rest_helper = mock_rest_helper

	var test_id int = 13
	resp, err := AncestryByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Activists"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}
}

func TestSuccessful_AsteroidBeltByID(t *testing.T) {
	jsonResponse := `{
		"name": "Inaro IX - Asteroid Belt 1",
		"position": {
		  "x": 809389301760,
		  "y": 151954759680,
		  "z": -221539000320
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	resp, err := AsteroidBeltByID(&test_id)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var resp_name string = "Inaro IX - Asteroid Belt 1"
	if *resp.Name != resp_name {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilID_AsteroidBeltByID(t *testing.T) {
	jsonResponse := `{
		"name": "Inaro IX - Asteroid Belt 1",
		"position": {
		  "x": 809389301760,
		  "y": 151954759680,
		  "z": -221539000320
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id *int = nil

	_, err := AsteroidBeltByID(test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "nil id" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCall_AsteroidBeltByID(t *testing.T) {
	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	_, err := AsteroidBeltByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshal_AsteroidBeltByID(t *testing.T) {
	jsonResponse := `{{
		"name": "Inaro IX - Asteroid Belt 1",
		"position": {
		  "x": 809389301760,
		  "y": 151954759680,
		  "z": -221539000320
		},
		"system_id": 30002788
	  }`

	b := []byte(jsonResponse)

	mock_rest_helper := &MockRestHelper{
		MockMakeCachingRESTCall: func(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	rest_helper = mock_rest_helper

	var test_id int = 1

	_, err := AsteroidBeltByID(&test_id)
	if err == nil {
		t.Error("Error is nil")
	}

}

type MockAddToRedisCacheType func(key string, value []byte, ttl int64)
type MockCheckRedisCacheType func(key string) (bool, []byte)

type MockRedisClient struct {
	MockAdd   MockAddToRedisCacheType
	MockCheck MockCheckRedisCacheType
}

func (m *MockRedisClient) AddToRedisCache(key string, value []byte, ttl int64) {
	m.MockAdd(key, value, ttl)
}

func (m *MockRedisClient) CheckRedisCache(key string) (bool, []byte) {
	return m.MockCheck(key)
}

type MockMakeCachingRESTCallType func(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error)

type MockRestHelper struct {
	MockMakeCachingRESTCall MockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
	return m.MockMakeCachingRESTCall(base_url, verb, body, additional_query_params, redis_query_key)
}
