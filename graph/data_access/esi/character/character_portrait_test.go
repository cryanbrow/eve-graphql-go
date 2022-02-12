package character

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	"github.com/cryanbrow/eve-graphql-go/graph/helpers"
)

func TestSuccessfulCharacterPortraitByID(t *testing.T) {
	jsonResponse := `{
		"px128x128": "https://images.evetech.net/characters/3018996/portrait?tenant=tranquility&size=128",
		"px256x256": "https://images.evetech.net/characters/3018996/portrait?tenant=tranquility&size=256",
		"px512x512": "https://images.evetech.net/characters/3018996/portrait?tenant=tranquility&size=512",
		"px64x64": "https://images.evetech.net/characters/3018996/portrait?tenant=tranquility&size=64"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	resp, err := CharacterPortraitByID(&testId, context.Background())
	if err != nil {
		t.Errorf("Error was not nil, %v", err)
	}
	var responseName string = "https://images.evetech.net/characters/3018996/portrait?tenant=tranquility&size=128"
	if *resp.Px128x128 != responseName {
		t.Errorf("Response was not as expected")
	}

}

func TestFailNilIDCharacterPortraitByID(t *testing.T) {
	jsonResponse := `{
		"px128x128": "https://images.evetech.net/characters/3018996/portrait?tenant=tranquility&size=128",
		"px256x256": "https://images.evetech.net/characters/3018996/portrait?tenant=tranquility&size=256",
		"px512x512": "https://images.evetech.net/characters/3018996/portrait?tenant=tranquility&size=512",
		"px64x64": "https://images.evetech.net/characters/3018996/portrait?tenant=tranquility&size=64"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId *int = nil

	_, err := CharacterPortraitByID(testId, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != helpers.NilId {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailRestCallCharacterPortraitByID(t *testing.T) {
	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return nil, nil, errors.New("failure")
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	_, err := CharacterPortraitByID(&testId, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	} else if err.Error() != "failure" {
		t.Errorf("Wrong error text: %s", err.Error())
	}

}

func TestFailUnmarshalCharacterPortraitByID(t *testing.T) {
	jsonResponse := `{{
		"px128x128": "https://images.evetech.net/characters/3018996/portrait?tenant=tranquility&size=128",
		"px256x256": "https://images.evetech.net/characters/3018996/portrait?tenant=tranquility&size=256",
		"px512x512": "https://images.evetech.net/characters/3018996/portrait?tenant=tranquility&size=512",
		"px64x64": "https://images.evetech.net/characters/3018996/portrait?tenant=tranquility&size=64"
	  }`

	b := []byte(jsonResponse)

	mockRestHelper := &MockRestHelper{
		CharacterMockMakeCachingRESTCall: func(baseUrl string, verb string, body bytes.Buffer, additionalQueryParams []configuration.Key_value, redisQueryKey string, ctx context.Context) ([]byte, http.Header, error) {
			return b, nil, nil
		},
	}
	restHelper = mockRestHelper

	var testId int = 1

	_, err := CharacterPortraitByID(&testId, context.Background())
	if err == nil {
		t.Error(helpers.NilError)
	}

}