package eve_character

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

func TestSuccessfulAttributesByID(t *testing.T) {
	jsonResponse := `{
		"bonus_remaps": 2,
		"charisma": 24,
		"intelligence": 25,
		"memory": 25,
		"perception": 25,
		"willpower": 25
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	resp, err := AttributesByID(context.Background(), &testID)
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var bonus_remaps = 2
	var charisma = 24
	var intelligence = 25
	var memory = 25
	var perception = 25
	var willpower = 25
	if *resp.BonusRemaps != bonus_remaps && *resp.Charisma != charisma && *resp.Intelligence != intelligence && *resp.Memory != memory && *resp.Perception != perception && *resp.Willpower != willpower {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDAttributesByID(t *testing.T) {
	jsonResponse := `{
		"bonus_remaps": 2,
		"charisma": 24,
		"intelligence": 25,
		"memory": 25,
		"perception": 25,
		"willpower": 25
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID *int

	_, err := AttributesByID(context.Background(), testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilID {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCallAttributesByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := AttributesByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshalAttributesByID(t *testing.T) {
	jsonResponse := `{{
		"bonus_remaps": 2,
		"charisma": 24,
		"intelligence": 25,
		"memory": 25,
		"perception": 25,
		"willpower": 25
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(ctx context.Context, baseURL string, verb string, body bytes.Buffer, additionalQueryParams []configuration.KeyValue, redisQueryKey string) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testID = 1

	_, err := AttributesByID(context.Background(), &testID)
	if err == nil {
		t.Error(helpers.NilError)
	}

}
