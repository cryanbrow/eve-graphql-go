package universe

import (
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

/***************************************
*             FactionByID              *
***************************************/

func TestSuccessfulInCacheFactionByID(t *testing.T) {
	jsonResponse := `{
		"corporation_id": 1000084,
		"description": "The largest of the five main empires, the Amarr Empire is a sprawling patch-work of feudal-like provinces held together by the might of the emperor. Religion has always played a big part in Amarrian politics and the Amarrians believe they are the rightful masters of the world, souring their relations with their neighbours. Another source of ill-feelings on part of the other empires is the fact that the Amarrians embrace slavery.",
		"faction_id": 500003,
		"is_unique": true,
		"militia_corporation_id": 1000179,
		"name": "Amarr Empire",
		"size_factor": 5,
		"solar_system_id": 30002187,
		"station_count": 1031,
		"station_system_count": 508
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

	var testId int = 500003
	resp, err := FactionByID(&testId)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Amarr Empire"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulNotInCacheFactionByID(t *testing.T) {
	var ancestriesJsonResponse string = `[
		{
			"corporation_id": 1000084,
			"description": "The largest of the five main empires, the Amarr Empire is a sprawling patch-work of feudal-like provinces held together by the might of the emperor. Religion has always played a big part in Amarrian politics and the Amarrians believe they are the rightful masters of the world, souring their relations with their neighbours. Another source of ill-feelings on part of the other empires is the fact that the Amarrians embrace slavery.",
			"faction_id": 500003,
			"is_unique": true,
			"militia_corporation_id": 1000179,
			"name": "Amarr Empire",
			"size_factor": 5,
			"solar_system_id": 30002187,
			"station_count": 1031,
			"station_system_count": 508
		  }
  ]`
	shouldReturn := setupNotInCacheRedis(ancestriesJsonResponse)
	if shouldReturn {
		return
	}

	var testId int = 500003
	resp, err := FactionByID(&testId)
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Amarr Empire"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestFailNilIDFactionByID(t *testing.T) {
	var testId *int = nil
	_, err := FactionByID(testId)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailUnmarshalInCacheFactionByID(t *testing.T) {
	jsonResponse := `{{
		"corporation_id": 1000084,
		"description": "The largest of the five main empires, the Amarr Empire is a sprawling patch-work of feudal-like provinces held together by the might of the emperor. Religion has always played a big part in Amarrian politics and the Amarrians believe they are the rightful masters of the world, souring their relations with their neighbours. Another source of ill-feelings on part of the other empires is the fact that the Amarrians embrace slavery.",
		"faction_id": 500003,
		"is_unique": true,
		"militia_corporation_id": 1000179,
		"name": "Amarr Empire",
		"size_factor": 5,
		"solar_system_id": 30002187,
		"station_count": 1031,
		"station_system_count": 508
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

	var testId int = 500003
	_, err := FactionByID(&testId)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailUnmarshalNotInCacheFactionByID(t *testing.T) {
	var ancestriesJsonResponse string = `[
	{{
		"corporation_id": 1000084,
		"description": "The largest of the five main empires, the Amarr Empire is a sprawling patch-work of feudal-like provinces held together by the might of the emperor. Religion has always played a big part in Amarrian politics and the Amarrians believe they are the rightful masters of the world, souring their relations with their neighbours. Another source of ill-feelings on part of the other empires is the fact that the Amarrians embrace slavery.",
		"faction_id": 500003,
		"is_unique": true,
		"militia_corporation_id": 1000179,
		"name": "Amarr Empire",
		"size_factor": 5,
		"solar_system_id": 30002187,
		"station_count": 1031,
		"station_system_count": 508
	  }
  ]`
	shouldReturn := setupNotInCacheRedis(ancestriesJsonResponse)
	if shouldReturn {
		return
	}

	var testId int = 500003
	_, err := FactionByID(&testId)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailRestNotInCacheFactionByID(t *testing.T) {
	shouldReturn := setupRESTFailureNotInCache()
	if shouldReturn {
		return
	}

	var testId int = 500003
	_, err := FactionByID(&testId)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}