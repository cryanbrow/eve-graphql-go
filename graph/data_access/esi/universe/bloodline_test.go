package universe

import (
	"context"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

/***************************************
*          BloodlineBeltByID           *
***************************************/

func TestSuccessfulInCacheBloodlineByID(t *testing.T) {
	jsonResponse := `{
		"bloodline_id": 5,
		"charisma": 3,
		"corporation_id": 1000066,
		"description": "True Amarrians are proud and supercilious, with a great sense of tradition and ancestry. They are considered arrogant and tyrannical by most others. The Empire's defeat at the hands of the mysterious Jovians, and the Minmatar uprising that followed, left an indelible mark on Amarrian culture. This double failure, a turning point in their history, has shaped an entire generation of policy and philosophy among the imperial elite.",
		"intelligence": 7,
		"memory": 6,
		"name": "Amarr",
		"perception": 4,
		"race_id": 4,
		"ship_type_id": 596,
		"willpower": 10
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
	resp, err := BloodlineByID(&testId, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Amarr"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestSuccessfulNotInCacheBloodlineByID(t *testing.T) {
	var ancestriesJsonResponse string = `[
		{
			"bloodline_id": 5,
			"charisma": 3,
			"corporation_id": 1000066,
			"description": "True Amarrians are proud and supercilious, with a great sense of tradition and ancestry. They are considered arrogant and tyrannical by most others. The Empire's defeat at the hands of the mysterious Jovians, and the Minmatar uprising that followed, left an indelible mark on Amarrian culture. This double failure, a turning point in their history, has shaped an entire generation of policy and philosophy among the imperial elite.",
			"intelligence": 7,
			"memory": 6,
			"name": "Amarr",
			"perception": 4,
			"race_id": 4,
			"ship_type_id": 596,
			"willpower": 10
		  }
  ]`
	shouldReturn := setupNotInCacheRedis(ancestriesJsonResponse)
	if shouldReturn {
		return
	}

	var testId int = 5
	resp, err := BloodlineByID(&testId, context.Background())
	if err != nil {
		t.Errorf(helpers.ErrorWasNotNil, err)
	}
	var responseName string = "Amarr"
	if *resp.Name != responseName {
		t.Errorf(helpers.ResponseWasNotAsExpected)
	}
}

func TestFailNilIDBloodlineByID(t *testing.T) {
	var testId *int = nil
	_, err := BloodlineByID(testId, context.Background())
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailUnmarshalInCacheBloodlineByID(t *testing.T) {
	jsonResponse := `{{
		"bloodline_id": 5,
		"charisma": 3,
		"corporation_id": 1000066,
		"description": "True Amarrians are proud and supercilious, with a great sense of tradition and ancestry. They are considered arrogant and tyrannical by most others. The Empire's defeat at the hands of the mysterious Jovians, and the Minmatar uprising that followed, left an indelible mark on Amarrian culture. This double failure, a turning point in their history, has shaped an entire generation of policy and philosophy among the imperial elite.",
		"intelligence": 7,
		"memory": 6,
		"name": "Amarr",
		"perception": 4,
		"race_id": 4,
		"ship_type_id": 596,
		"willpower": 10
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

	var testId int = 5
	_, err := BloodlineByID(&testId, context.Background())
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailUnmarshalNotInCacheBloodlineByID(t *testing.T) {
	var ancestriesJsonResponse string = `[
	{{
		"bloodline_id": 5,
		"charisma": 3,
		"corporation_id": 1000066,
		"description": "True Amarrians are proud and supercilious, with a great sense of tradition and ancestry. They are considered arrogant and tyrannical by most others. The Empire's defeat at the hands of the mysterious Jovians, and the Minmatar uprising that followed, left an indelible mark on Amarrian culture. This double failure, a turning point in their history, has shaped an entire generation of policy and philosophy among the imperial elite.",
		"intelligence": 7,
		"memory": 6,
		"name": "Amarr",
		"perception": 4,
		"race_id": 4,
		"ship_type_id": 596,
		"willpower": 10
	  }
  ]`
	shouldReturn := setupNotInCacheRedis(ancestriesJsonResponse)
	if shouldReturn {
		return
	}

	var testId int = 5
	_, err := BloodlineByID(&testId, context.Background())
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailRestNotInCacheBloodlineByID(t *testing.T) {
	shouldReturn := setupRESTFailureNotInCache()
	if shouldReturn {
		return
	}

	var testId int = 5
	_, err := BloodlineByID(&testId, context.Background())
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}
