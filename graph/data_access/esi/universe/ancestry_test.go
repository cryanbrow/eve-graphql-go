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
		MockAdd: func(ctx context.Context, key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(ctx context.Context, key string) (bool, []byte) {
			return true, b
		},
	}
	CachingClient = mockCachingClient

	var testID = 1
	resp, err := AncestryByID(context.Background(), &testID)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Activists"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulNotInCacheAncestryByID(t *testing.T) {
	var ancestriesJSONResponse = `[
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
	shouldReturn := setupNotInCacheRedis(ancestriesJSONResponse)
	if shouldReturn {
		return
	}

	var testID = 13
	resp, err := AncestryByID(context.Background(), &testID)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Activists"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func setupNotInCacheRedis(jsonResponse string) bool {
	b := []byte(jsonResponse)
	mockCachingClient := &MockCachingClient{
		MockAdd: func(ctx context.Context, key string, value []byte, ttl int64) {
			//This method does nothing when mocked
		},
		MockCheck: func(ctx context.Context, key string) (bool, []byte) {
			return false, nil
		},
	}
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}

	CachingClient = mockCachingClient
	restHelper = mockRestHelper
	return false
}

func TestFailNilIDAncestryByID(t *testing.T) {
	var testID *int
	_, err := AncestryByID(context.Background(), testID)
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
		MockAdd: func(ctx context.Context, key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(ctx context.Context, key string) (bool, []byte) {
			return true, b
		},
	}
	CachingClient = mockCachingClient

	var testID = 13
	_, err := AncestryByID(context.Background(), &testID)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailUnmarshalNotInCacheAncestryByID(t *testing.T) {
	var ancestriesJSONResponse = `[
	{{
	  "bloodline_id": 7,
	  "description": "The Gallente prize political activism more so than other Empires. Many devote their efforts towards one or more causes that suit their ambitions. Activists understand that things will never change for the better unless someone has the courage to fight the good fight.",
	  "icon_id": 1653,
	  "id": 13,
	  "name": "Activists",
	  "short_description": "Making the universe a better place, one fight at a time."
	}
  ]`
	shouldReturn := setupNotInCacheRedis(ancestriesJSONResponse)
	if shouldReturn {
		return
	}

	var testID = 13
	_, err := AncestryByID(context.Background(), &testID)
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

	var testID = 13
	_, err := AncestryByID(context.Background(), &testID)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func setupRESTFailureNotInCache() bool {
	mockCachingClient := &MockCachingClient{
		MockAdd: func(ctx context.Context, key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(ctx context.Context, key string) (bool, []byte) {
			return false, nil
		},
	}
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KevValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}

	CachingClient = mockCachingClient
	restHelper = mockRestHelper
	return false
}
