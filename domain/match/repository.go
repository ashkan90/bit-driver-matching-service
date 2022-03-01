package match

import (
	"bit-driver-matching-service/adapters/http_client"
	"bit-driver-matching-service/config"
	"bit-driver-matching-service/request"
	"bit-driver-matching-service/response"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/mercari/go-circuitbreaker"
	"log"
	"net/http"
)

type RepositoryMatch struct {
	MatchServiceConfig *config.Service
	breaker            *circuitbreaker.CircuitBreaker
	client             http_client.ClientImplementation
}

func NewRepositoryMatch(config *config.Service, breaker *circuitbreaker.CircuitBreaker) *RepositoryMatch {
	return &RepositoryMatch{
		MatchServiceConfig: config,
		breaker:            breaker,
		client:             http_client.NewClient(),
	}
}

func (r *RepositoryMatch) FindNearest(loc request.CustomerLocation) response.NearestDriver {
	var req, _ = json.Marshal(&loc)
	var uri = r.MatchServiceConfig.URL + r.MatchServiceConfig.Path

	var resp, err = r.breaker.Do(context.Background(), func() (interface{}, error) {
		var resp, err = r.client.Do(http.MethodGet, uri, bytes.NewReader(req))
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, errors.New("status is not ok")
		}

		return resp, nil
	})

	if err != nil {
		log.Println("an error has occurred: ", err)
		return response.NearestDriver{}
	}

	var driver response.NearestDriver
	_ = json.NewDecoder(resp.(*http.Response).Body).Decode(&driver)

	return driver
}
