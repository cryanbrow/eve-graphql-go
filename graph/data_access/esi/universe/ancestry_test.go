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
*             AncestryByID             *
***************************************/

func TestSuccessfulInCacheAncestryByID(t *testing.T) {
	jsonResponse := `{
		"bloodline_id": 7,
		"description": "The Gallente prize political activism more so than other Empires. Many devote their efforts towards one or more causes that suit their ambitions. Activists understand that things will never change for the better unless someone has the courage to fight the good fight.",
		"icon_id": 1653,
		"id": 13,
		"name": "Activists",
		"short_description": "Making the universe a better place, one fight at a time."
	  }`
	b := []byte(jsonResponse)

	mockCachingClient := &MockCachingClient{
		MockAdd: func(key string, value []byte, ttl int64, ctx context.Context) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string, ctx context.Context) (bool, []byte) {
			return true, b
		},
	}
	CachingClient = mockCachingClient

	var testId int = 1
	resp, err := AncestryByID(&testId, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Activists"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulNotInCacheAncestryByID(t *testing.T) {
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
	//Method returns nothing so needs no implementation
	shouldReturn := setupNotInCacheRedis(ancestriesJsonResponse)
	if shouldReturn {
		return
	}

	var testId int = 13
	resp, err := AncestryByID(&testId, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Activists"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func setupNotInCacheRedis(jsonResponse string) bool {
	b := []byte(jsonResponse)
	mockCachingClient := &MockCachingClient{
		MockAdd: func(key string, value []byte, ttl int64, ctx context.Context) {
			//This method does nothing when mocked
		},
		MockCheck: func(key string, ctx context.Context) (bool, []byte) {
			return false, nil
		},
	}
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}

	CachingClient = mockCachingClient
	restHelper = mockRestHelper
	return false
}

func TestFailNilIDAncestryByID(t *testing.T) {
	var testId *int = nil
	_, err := AncestryByID(testId, context.Background())
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailUnmarshalInCacheAncestryByID(t *testing.T) {
	jsonResponse := `{{
		"bloodline_id": 7,
		"description": "The Gallente prize political activism more so than other Empires. Many devote their efforts towards one or more causes that suit their ambitions. Activists understand that things will never change for the better unless someone has the courage to fight the good fight.",
		"icon_id": 1653,
		"id": 13,
		"name": "Activists",
		"short_description": "Making the universe a better place, one fight at a time."
	  }`
	b := []byte(jsonResponse)

	mockCachingClient := &MockCachingClient{
		MockAdd: func(key string, value []byte, ttl int64, ctx context.Context) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string, ctx context.Context) (bool, []byte) {
			return true, b
		},
	}
	CachingClient = mockCachingClient

	var testId int = 13
	_, err := AncestryByID(&testId, context.Background())
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailUnmarshalNotInCacheAncestryByID(t *testing.T) {
	var ancestriesJsonResponse string = `[
	{{
	  "bloodline_id": 7,
	  "description": "The Gallente prize political activism more so than other Empires. Many devote their efforts towards one or more causes that suit their ambitions. Activists understand that things will never change for the better unless someone has the courage to fight the good fight.",
	  "icon_id": 1653,
	  "id": 13,
	  "name": "Activists",
	  "short_description": "Making the universe a better place, one fight at a time."
	}
  ]`
	shouldReturn := setupNotInCacheRedis(ancestriesJsonResponse)
	if shouldReturn {
		return
	}

	var testId int = 13
	_, err := AncestryByID(&testId, context.Background())
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailRestNotInCacheAncestryByID(t *testing.T) {
	//Method returns nothing so needs no implementation
	shouldReturn := setupRESTFailureNotInCache()
	if shouldReturn {
		return
	}

	var testId int = 13
	_, err := AncestryByID(&testId, context.Background())
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func setupRESTFailureNotInCache() bool {
	mockCachingClient := &MockCachingClient{
		MockAdd: func(key string, value []byte, ttl int64, ctx context.Context) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string, ctx context.Context) (bool, []byte) {
			return false, nil
		},
	}
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}

	CachingClient = mockCachingClient
	restHelper = mockRestHelper
	return false
}
