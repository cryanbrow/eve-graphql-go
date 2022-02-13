package universe

import (
	"context"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

/***************************************
*               RaceByID               *
***************************************/

func TestSuccessfulInCacheRaceByID(t *testing.T) {
	jsonResponse := `{
		"alliance_id": 500001,
		"description": "Founded on the tenets of patriotism and hard work that carried its ancestors through hardships on an inhospitable homeworld, the Caldari State is today a corporate dictatorship, led by rulers who are determined to see it return to the meritocratic ideals of old. Ruthless and efficient in the boardroom as well as on the battlefield, the Caldari are living emblems of strength, persistence, and dignity.",
		"name": "Caldari",
		"race_id": 1
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

	var testID = 1
	resp, err := RaceByID(context.Background(), &testID)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Caldari"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulNotInCacheRaceByID(t *testing.T) {
	var ancestriesJSONResponse string = `[{
		"alliance_id": 500001,
		"description": "Founded on the tenets of patriotism and hard work that carried its ancestors through hardships on an inhospitable homeworld, the Caldari State is today a corporate dictatorship, led by rulers who are determined to see it return to the meritocratic ideals of old. Ruthless and efficient in the boardroom as well as on the battlefield, the Caldari are living emblems of strength, persistence, and dignity.",
		"name": "Caldari",
		"race_id": 1
	  }]`
	//Method returns nothing so needs no implementation
	shouldReturn := setupNotInCacheRedis(ancestriesJSONResponse)
	if shouldReturn {
		return
	}

	var testID = 1
	resp, err := RaceByID(context.Background(), &testID)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName = "Caldari"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestFailNilIDRaceByID(t *testing.T) {
	var testID *int
	_, err := RaceByID(context.Background(), testID)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailUnmarshalInCacheRaceByID(t *testing.T) {
	jsonResponse := `{{
		"alliance_id": 500001,
		"description": "Founded on the tenets of patriotism and hard work that carried its ancestors through hardships on an inhospitable homeworld, the Caldari State is today a corporate dictatorship, led by rulers who are determined to see it return to the meritocratic ideals of old. Ruthless and efficient in the boardroom as well as on the battlefield, the Caldari are living emblems of strength, persistence, and dignity.",
		"name": "Caldari",
		"race_id": 1
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

	var testID = 1
	_, err := RaceByID(context.Background(), &testID)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailUnmarshalNotInCacheRaceByID(t *testing.T) {
	var ancestriesJSONResponse string = `[
	{{
		"alliance_id": 500001,
		"description": "Founded on the tenets of patriotism and hard work that carried its ancestors through hardships on an inhospitable homeworld, the Caldari State is today a corporate dictatorship, led by rulers who are determined to see it return to the meritocratic ideals of old. Ruthless and efficient in the boardroom as well as on the battlefield, the Caldari are living emblems of strength, persistence, and dignity.",
		"name": "Caldari",
		"race_id": 1
	  }
  ]`
	shouldReturn := setupNotInCacheRedis(ancestriesJSONResponse)
	if shouldReturn {
		return
	}

	var testID = 1
	_, err := RaceByID(context.Background(), &testID)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailRestNotInCacheRaceByID(t *testing.T) {
	//Method returns nothing so needs no implementation
	shouldReturn := setupRESTFailureNotInCache()
	if shouldReturn {
		return
	}

	var testID = 1
	_, err := RaceByID(context.Background(), &testID)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}
