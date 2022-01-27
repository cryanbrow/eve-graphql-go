package helpers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	cache "github.com/cryanbrow/eve-graphql-go/graph/caching"
	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	log "github.com/sirupsen/logrus"
)

func MakeCachingRESTCall(base_url string, verb string, body bytes.Buffer, additional_query_params []configuration.Key_value, redis_query_key string) ([]byte, http.Header, error) {
	inCache, result := cache.CheckRedisCache(redis_query_key)
	if !inCache {
		crest_url, err := url.Parse(base_url)
		if err != nil {
			log.WithFields(log.Fields{"base_url": base_url, "verb": verb}).Errorf("Failed to Parse URL with Error : %v", err)
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
			return make([]byte, 0), nil, err
		}
		if err != nil {
			log.WithFields(log.Fields{"url": url}).Errorf("Could not read response for body. : %v", err)
			return make([]byte, 0), nil, err
		}
		cache.AddToRedisCache(redis_query_key, responseBytes, ESI_ttl_to_millis(h.Get("expires")))
		return responseBytes, h, nil
	}
	return result, nil, nil
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	Client = &http.Client{}
}
