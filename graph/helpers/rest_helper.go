package helpers

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	cache "github.com/cryanbrow/eve-graphql-go/graph/caching"
	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	log "github.com/sirupsen/logrus"
)

type RestHelperClient struct {
}

func (r *RestHelperClient) MakeCachingRESTCall(baseUrl string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
	inCache, result := Redis_client.CheckRedisCache(redis_query_key)
	if !inCache {
		crest_url, err := url.Parse(baseUrl)
		if err != nil {
			log.WithFields(log.Fields{"baseUrl": baseUrl, "verb": verb}).Errorf("Failed to Parse URL with Error : %v", err)
			return nil, nil, err
		}
		queryParameters := crest_url.Query()
		for _, kv := range configuration.AppConfig.Esi.Default.Query_params {
			queryParameters.Add(kv.Key, kv.Value)
		}
		for _, kv := range additional_query_params {
			queryParameters.Add(kv.Key, kv.Value)
		}

		crest_url.RawQuery = queryParameters.Encode()
		url := crest_url.String()

		log.WithFields(log.Fields{"url": url}).Info("Making REST Call")
		request, err := http.NewRequest(verb, url, &body)
		if err != nil {
			log.WithFields(log.Fields{"url": url}).Errorf("Could not build request. : %v", err)
			return make([]byte, 0), nil, err
		}
		response, err := Client.Do(request)
		if err != nil {
			log.WithFields(log.Fields{"url": url}).Errorf("Could not make request. : %v", err)
			return make([]byte, 0), nil, err
		}

		h := response.Header
		responseBytes, err := ioutil.ReadAll(response.Body)
		if response.StatusCode != 200 {
			log.WithFields(log.Fields{"url": url, "status_code": response.StatusCode}).Errorf("Received bad status code. : %v", err)
			return make([]byte, 0), nil, errors.New(response.Status)
		}
		if err != nil {
			log.WithFields(log.Fields{"url": url}).Errorf("Could not read response for body. : %v", err)
			return make([]byte, 0), nil, err
		}
		Redis_client.AddToRedisCache(redis_query_key, responseBytes, ESI_ttl_to_millis(h.Get("expires")))
		return responseBytes, h, nil
	}
	return result, nil, nil
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RedisClient interface {
	AddToRedisCache(key string, value []byte, ttl int64)
	CheckRedisCache(key string) (bool, []byte)
}

var (
	Client       HTTPClient
	Redis_client RedisClient
)

func SetupRestHelper() {
	Client = &http.Client{}
	Redis_client = &cache.Client{}
}
