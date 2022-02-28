package match

import (
	"bit-driver-matching-service/adapters/http_client"
	"bit-driver-matching-service/config"
	"bit-driver-matching-service/request"
	"bit-driver-matching-service/response"
	"bytes"
	"encoding/json"
	"net/http"
)

type RepositoryMatch struct {
	MatchServiceConfig *config.Service
}

func NewRepositoryMatch(config *config.Service) *RepositoryMatch {
	return &RepositoryMatch{
		MatchServiceConfig: config,
	}
}

func (r *RepositoryMatch) FindNearest(loc request.CustomerLocation) response.NearestDriver {
	var req, _ = json.Marshal(&loc)

	var uri = r.MatchServiceConfig.URL + "/nearest-driver-location"
	var resp, err = http_client.NewClient().Do(http.MethodGet, uri, bytes.NewReader(req))
	if err != nil {
		return response.NearestDriver{}
	}

	defer resp.Body.Close()

	var driver response.NearestDriver
	json.NewDecoder(resp.Body).Decode(&driver)

	return driver
}
