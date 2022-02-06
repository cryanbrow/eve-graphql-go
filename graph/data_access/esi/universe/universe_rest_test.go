package universe

import (
	"bytes"
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
	resp, err := AncestryByID(&testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Activists"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
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
	b := []byte(ancestriesJsonResponse)
	mockRedisClient := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}

	RedisClient = mockRedisClient
	restHelper = mockRestHelper

	var testId int = 13
	resp, err := AncestryByID(&testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Activists"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
	}
}

func TestFailNilIDAncestryByID(t *testing.T) {
	var testId *int = nil
	_, err := AncestryByID(testId)
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

	mockRedisClient := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string) (bool, []byte) {
			return true, b
		},
	}
	RedisClient = mockRedisClient

	var testId int = 13
	_, err := AncestryByID(&testId)
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
	b := []byte(ancestriesJsonResponse)
	mockRedisClient := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	RedisClient = mockRedisClient

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 13
	_, err := AncestryByID(&testId)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailRestNotInCacheAncestryByID(t *testing.T) {
	mockRedisClient := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}

	RedisClient = mockRedisClient
	restHelper = mockRestHelper

	var testId int = 13
	_, err := AncestryByID(&testId)
	if err == nil {
		t.Errorf("Error was nil")
	}
}

/***************************************
*          AsteroidBeltDetails         *
***************************************/

func TestSuccessfulAsteroidBeltDetails(t *testing.T) {
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

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = new(int)
	*testId = 1
	var ids []*int = make([]*int, 1)
	ids[0] = testId

	resp, err := AsteroidBeltDetails(ids)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Inaro IX - Asteroid Belt 1"
	if *resp[0].Name != responseName {
		t.Errorf("Response was not as expected")
	}
}

func TestFailNilIDAsteroidBeltDetails(t *testing.T) {
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

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = new(int)
	*testId = 1
	var ids []*int = make([]*int, 2)
	ids[0] = testId

	_, err := AsteroidBeltDetails(ids)
	if err == nil {
		t.Errorf("Error was nil")
	}
}

/***************************************
*           AsteroidBeltByID           *
***************************************/

func TestSuccessfulAsteroidBeltByID(t *testing.T) {
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

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	resp, err := AsteroidBeltByID(&testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Inaro IX - Asteroid Belt 1"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDAsteroidBeltByID(t *testing.T) {
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

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = nil

	_, err := AsteroidBeltByID(testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCallAsteroidBeltByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	_, err := AsteroidBeltByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshalAsteroidBeltByID(t *testing.T) {
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

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	_, err := AsteroidBeltByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

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
	resp, err := BloodlineByID(&testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Amarr"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
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
	b := []byte(ancestriesJsonResponse)
	mockRedisClient := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation,
		},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}

	RedisClient = mockRedisClient
	restHelper = mockRestHelper

	var testId int = 5
	resp, err := BloodlineByID(&testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Amarr"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
	}
}

func TestFailNilIDBloodlineByID(t *testing.T) {
	var testId *int = nil
	_, err := BloodlineByID(testId)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailUnmarshalInCacheBBloodlineByID(t *testing.T) {
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
	_, err := BloodlineByID(&testId)
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
	b := []byte(ancestriesJsonResponse)
	mockRedisClient := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	RedisClient = mockRedisClient

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 5
	_, err := BloodlineByID(&testId)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailRestNotInCacheBloodlineByID(t *testing.T) {
	mockRedisClient := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}

	RedisClient = mockRedisClient
	restHelper = mockRestHelper

	var testId int = 5
	_, err := BloodlineByID(&testId)
	if err == nil {
		t.Errorf("Error was nil")
	}
}

/***************************************
*             CategoryByID             *
***************************************/

func TestSuccessfulCategoryByID(t *testing.T) {
	jsonResponse := `{
		"category_id": 5,
		"groups": [
		  23,
		  24,
		  872,
		  876,
		  943,
		  1301,
		  1311,
		  1739,
		  1875
		],
		"name": "Accessories",
		"published": true
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 5

	resp, err := CategoryByID(&testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Accessories"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDCategoryByID(t *testing.T) {
	jsonResponse := `{
		"category_id": 5,
		"groups": [
		  23,
		  24,
		  872,
		  876,
		  943,
		  1301,
		  1311,
		  1739,
		  1875
		],
		"name": "Accessories",
		"published": true
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = nil

	_, err := CategoryByID(testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCallCategoryByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	_, err := CategoryByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshalCategoryByID(t *testing.T) {
	jsonResponse := `{{
		"category_id": 5,
		"groups": [
		  23,
		  24,
		  872,
		  876,
		  943,
		  1301,
		  1311,
		  1739,
		  1875
		],
		"name": "Accessories",
		"published": true
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 5

	_, err := CategoryByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

/***************************************
*          ConstellationsByIDs         *
***************************************/

func TestSuccessfulConstellationsByIDs(t *testing.T) {
	jsonResponse := `{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = new(int)
	*testId = 20000019
	var ids []*int = make([]*int, 1)
	ids[0] = testId

	resp, err := ConstellationsByIDs(ids)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Ihilakken"
	if *resp[0].Name != responseName {
		t.Errorf("Response was not as expected")
	}
}

func TestFailNilIDConstellationsByIDs(t *testing.T) {
	jsonResponse := `{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = new(int)
	*testId = 20000019
	var ids []*int = make([]*int, 2)
	ids[0] = testId

	_, err := ConstellationsByIDs(ids)
	if err == nil {
		t.Errorf("Error was nil")
	}
}

/***************************************
*          ConstellationByID           *
***************************************/

func TestSuccessfulConstellationByID(t *testing.T) {
	jsonResponse := `{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 20000019

	resp, err := ConstellationByID(&testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Ihilakken"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDConstellationByID(t *testing.T) {
	jsonResponse := `{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = nil

	_, err := ConstellationByID(testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCallConstellationByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 20000019

	_, err := ConstellationByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshalConstellationByID(t *testing.T) {
	jsonResponse := `{{
		"constellation_id": 20000019,
		"name": "Ihilakken",
		"position": {
		  "x": -143645654698282130,
		  "y": 52909580254258400,
		  "z": 109619376865938180
		},
		"region_id": 10000002,
		"systems": [
		  30000132,
		  30000133,
		  30000134,
		  30000135,
		  30000136,
		  30000137,
		  30000138,
		  30021407
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 20000019

	_, err := ConstellationByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

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
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Amarr Empire"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
	}
}

func TestSuccessfulNotInCache_FactionByID(t *testing.T) {
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
	b := []byte(ancestriesJsonResponse)
	mockRedisClient := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}

	RedisClient = mockRedisClient
	restHelper = mockRestHelper

	var testId int = 500003
	resp, err := FactionByID(&testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Amarr Empire"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
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
	b := []byte(ancestriesJsonResponse)
	mockRedisClient := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	RedisClient = mockRedisClient

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 500003
	_, err := FactionByID(&testId)
	if err == nil {
		t.Errorf(helpers.NilError)
	}
}

func TestFailRestNotInCacheFactionByID(t *testing.T) {
	mockRedisClient := &MockRedisClient{
		MockAdd: func(key string, value []byte, ttl int64) {
			//Method returns nothing so needs no implementation
		},
		MockCheck: func(key string) (bool, []byte) {
			return false, nil
		},
	}
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}

	RedisClient = mockRedisClient
	restHelper = mockRestHelper

	var testId int = 500003
	_, err := FactionByID(&testId)
	if err == nil {
		t.Errorf("Error was nil")
	}
}

/***************************************
*             GraphicByID              *
***************************************/

func TestSuccessfulGraphicByID(t *testing.T) {
	jsonResponse := `{
		"graphic_file": "res:/dx9/model/Turret/Energy/Pulse/M/Pulse_Heavy_T1.red",
		"graphic_id": 21573,
		"sof_fation_name": "amarrnavy",
		"sof_race_name": "amarr"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 21573

	resp, err := GraphicByID(&testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "amarrnavy"
	if *resp.SofFationName != responseName {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDGraphicByID(t *testing.T) {
	jsonResponse := `{
		"graphic_file": "res:/dx9/model/Turret/Energy/Pulse/M/Pulse_Heavy_T1.red",
		"graphic_id": 21573,
		"sof_fation_name": "amarrnavy",
		"sof_race_name": "amarr"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = nil

	_, err := GraphicByID(testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCallGraphicByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 21573

	_, err := GraphicByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshalGraphicByID(t *testing.T) {
	jsonResponse := `{{
		"graphic_file": "res:/dx9/model/Turret/Energy/Pulse/M/Pulse_Heavy_T1.red",
		"graphic_id": 21573,
		"sof_fation_name": "amarrnavy",
		"sof_race_name": "amarr"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 21573

	_, err := GraphicByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

/***************************************
*              GroupByID               *
***************************************/

func TestSuccessfulGroupByID(t *testing.T) {
	jsonResponse := `{
		"category_id": 6,
		"group_id": 25,
		"name": "Frigate",
		"published": true,
		"types": [
		  582,
		  583,
		  584,
		  585,
		  586,
		  587,
		  589,
		  590,
		  591,
		  592,
		  593,
		  594,
		  595,
		  597,
		  598,
		  599,
		  600,
		  602,
		  603,
		  605,
		  607,
		  608,
		  609,
		  613,
		  614,
		  616,
		  618,
		  619,
		  1896,
		  1898,
		  1900,
		  1902,
		  2161,
		  3532,
		  3751,
		  3753,
		  3766,
		  3768,
		  11019,
		  11940,
		  11942,
		  17360,
		  17619,
		  17703,
		  17705,
		  17707,
		  17812,
		  17841,
		  17924,
		  17926,
		  17928,
		  17930,
		  17932,
		  29248,
		  32880,
		  32983,
		  32985,
		  32987,
		  32989,
		  33190,
		  33468,
		  33655,
		  33657,
		  33659,
		  33661,
		  33663,
		  33665,
		  33667,
		  33669,
		  33677,
		  33816,
		  34443,
		  37453,
		  37454,
		  37455,
		  37456,
		  47269,
		  54731,
		  58745
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 25

	resp, err := GroupByID(&testId)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "Frigate"
	if *resp.Name != responseName {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDGroupByID(t *testing.T) {
	jsonResponse := `{
		"category_id": 6,
		"group_id": 25,
		"name": "Frigate",
		"published": true,
		"types": [
		  582,
		  583,
		  584,
		  585,
		  586,
		  587,
		  589,
		  590,
		  591,
		  592,
		  593,
		  594,
		  595,
		  597,
		  598,
		  599,
		  600,
		  602,
		  603,
		  605,
		  607,
		  608,
		  609,
		  613,
		  614,
		  616,
		  618,
		  619,
		  1896,
		  1898,
		  1900,
		  1902,
		  2161,
		  3532,
		  3751,
		  3753,
		  3766,
		  3768,
		  11019,
		  11940,
		  11942,
		  17360,
		  17619,
		  17703,
		  17705,
		  17707,
		  17812,
		  17841,
		  17924,
		  17926,
		  17928,
		  17930,
		  17932,
		  29248,
		  32880,
		  32983,
		  32985,
		  32987,
		  32989,
		  33190,
		  33468,
		  33655,
		  33657,
		  33659,
		  33661,
		  33663,
		  33665,
		  33667,
		  33669,
		  33677,
		  33816,
		  34443,
		  37453,
		  37454,
		  37455,
		  37456,
		  47269,
		  54731,
		  58745
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = nil

	_, err := GroupByID(testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCallGroupByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 25

	_, err := GroupByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshalGroupByID(t *testing.T) {
	jsonResponse := `{{
		"category_id": 6,
		"group_id": 25,
		"name": "Frigate",
		"published": true,
		"types": [
		  582,
		  583,
		  584,
		  585,
		  586,
		  587,
		  589,
		  590,
		  591,
		  592,
		  593,
		  594,
		  595,
		  597,
		  598,
		  599,
		  600,
		  602,
		  603,
		  605,
		  607,
		  608,
		  609,
		  613,
		  614,
		  616,
		  618,
		  619,
		  1896,
		  1898,
		  1900,
		  1902,
		  2161,
		  3532,
		  3751,
		  3753,
		  3766,
		  3768,
		  11019,
		  11940,
		  11942,
		  17360,
		  17619,
		  17703,
		  17705,
		  17707,
		  17812,
		  17841,
		  17924,
		  17926,
		  17928,
		  17930,
		  17932,
		  29248,
		  32880,
		  32983,
		  32985,
		  32987,
		  32989,
		  33190,
		  33468,
		  33655,
		  33657,
		  33659,
		  33661,
		  33663,
		  33665,
		  33667,
		  33669,
		  33677,
		  33816,
		  34443,
		  37453,
		  37454,
		  37455,
		  37456,
		  47269,
		  54731,
		  58745
		]
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		MockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 25

	_, err := GroupByID(&testId)
	if err == nil {
		t.Error(helpers.NilError)
	}

}

/***************************************
*             MOCK SECTION             *
***************************************/

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

type MockMakeCachingRESTCallType func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error)

type MockRestHelper struct {
	MockMakeCachingRESTCall MockMakeCachingRESTCallType
}

func (m *MockRestHelper) MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string) ([]byte, http.Header, error) {
	return m.MockMakeCachingRESTCall(baseUrl, verb, body, additionalQueryParams, redisQueryKey)
}
