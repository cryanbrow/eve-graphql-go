package universe

import (
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

	mockRedisClient := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string) (bool, []byte) {
			return true, b
		},
	}
	RedisClient = mockRedisClient

	var testId int = 1
	resp, err := RaceByID(&testId)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Caldari"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulNotInCacheRaceByID(t *testing.T) {
	var ancestriesJsonResponse string = `[{
		"alliance_id": 500001,
		"description": "Founded on the tenets of patriotism and hard work that carried its ancestors through hardships on an inhospitable homeworld, the Caldari State is today a corporate dictatorship, led by rulers who are determined to see it return to the meritocratic ideals of old. Ruthless and efficient in the boardroom as well as on the battlefield, the Caldari are living emblems of strength, persistence, and dignity.",
		"name": "Caldari",
		"race_id": 1
	  }]`
	//Method returns nothing so needs no implementation
	shouldReturn := setupNotInCacheRedis(ancestriesJsonResponse)
	if shouldReturn {
		return
	}

	var testId int = 1
	resp, err := RaceByID(&testId)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Caldari"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestFailNilIDRaceByID(t *testing.T) {
	var testId *int = nil
	_, err := RaceByID(testId)
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

	mockRedisClient := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string) (bool, []byte) {
			return true, b
		},
	}
	RedisClient = mockRedisClient

	var testId int = 1
	_, err := RaceByID(&testId)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailUnmarshalNotInCacheRaceByID(t *testing.T) {
	var ancestriesJsonResponse string = `[
	{{
		"alliance_id": 500001,
		"description": "Founded on the tenets of patriotism and hard work that carried its ancestors through hardships on an inhospitable homeworld, the Caldari State is today a corporate dictatorship, led by rulers who are determined to see it return to the meritocratic ideals of old. Ruthless and efficient in the boardroom as well as on the battlefield, the Caldari are living emblems of strength, persistence, and dignity.",
		"name": "Caldari",
		"race_id": 1
	  }
  ]`
	shouldReturn := setupNotInCacheRedis(ancestriesJsonResponse)
	if shouldReturn {
		return
	}

	var testId int = 1
	_, err := RaceByID(&testId)
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

	var testId int = 1
	_, err := RaceByID(&testId)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}
