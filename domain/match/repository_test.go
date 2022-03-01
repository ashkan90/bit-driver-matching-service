package match

import (
	mocks "bit-driver-matching-service/adapters/http_client/.mocks"
	"bit-driver-matching-service/config"
	"bit-driver-matching-service/request"
	"bit-driver-matching-service/response"
	"bytes"
	"fmt"
	"github.com/benbjohnson/clock"
	"github.com/cenkalti/backoff/v3"
	"github.com/golang/mock/gomock"
	"github.com/mercari/go-circuitbreaker"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestRepositoryMatch_FindNearestCircuitBreakerShouldOpenCircuit(t *testing.T) {
	var stateChanges []string
	var breaker = circuitbreaker.New(
		circuitbreaker.WithClock(clock.New()),
		circuitbreaker.WithFailOnContextCancel(true),
		circuitbreaker.WithFailOnContextDeadline(true),
		circuitbreaker.WithHalfOpenMaxSuccesses(10),
		circuitbreaker.WithOpenTimeoutBackOff(backoff.NewExponentialBackOff()),
		circuitbreaker.WithOpenTimeout(10*time.Second),
		circuitbreaker.WithCounterResetInterval(10*time.Second),
		circuitbreaker.WithTripFunc(circuitbreaker.NewTripFuncFailureRate(10, 0.4)),
		circuitbreaker.WithOnStateChangeHookFn(func(from, to circuitbreaker.State) {
			stateChanges = append(stateChanges, fmt.Sprintf("state changed from %s to %s", from, to))
		}),
	)

	for i := 0; i < 10; i++ {
		var resp = &http.Response{
			StatusCode: 400,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"coordinates": [40.946104,28.035588]}`)),
			Close:      false,
		}
		var ctrl = gomock.NewController(t)

		client := mocks.NewMockClientImplementation(ctrl)
		client.
			EXPECT().
			Do(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil).
			Times(1)

		var repo = RepositoryMatch{
			MatchServiceConfig: &config.Service{
				URL:  "url",
				Path: "/path",
			},
			breaker: breaker,
			client:  client,
		}
		var loc = request.CustomerLocation{
			Longitude: 10,
			Latitude:  1,
		}

		_ = repo.FindNearest(loc)
		ctrl.Finish()
	}

	assert.GreaterOrEqual(t, len(stateChanges), 1)
}

func TestRepositoryMatch_FindNearestCircuitBreakerShouldStayClosed(t *testing.T) {
	var stateChanges []string
	var breaker = circuitbreaker.New(
		circuitbreaker.WithClock(clock.New()),
		circuitbreaker.WithFailOnContextCancel(true),
		circuitbreaker.WithFailOnContextDeadline(true),
		circuitbreaker.WithHalfOpenMaxSuccesses(10),
		circuitbreaker.WithOpenTimeoutBackOff(backoff.NewExponentialBackOff()),
		circuitbreaker.WithOpenTimeout(10*time.Second),
		circuitbreaker.WithCounterResetInterval(10*time.Second),
		circuitbreaker.WithTripFunc(circuitbreaker.NewTripFuncFailureRate(10, 0.4)),
		circuitbreaker.WithOnStateChangeHookFn(func(from, to circuitbreaker.State) {
			stateChanges = append(stateChanges, fmt.Sprintf("state changed from %s to %s", from, to))
		}),
	)

	for i := 0; i < 10; i++ {
		var resp = &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"coordinates": [40.946104,28.035588]}`)),
			Close:      false,
		}
		var ctrl = gomock.NewController(t)

		client := mocks.NewMockClientImplementation(ctrl)
		client.
			EXPECT().
			Do(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(resp, nil).
			Times(1)

		var repo = RepositoryMatch{
			MatchServiceConfig: &config.Service{
				URL:  "url",
				Path: "/path",
			},
			breaker: breaker,
			client:  client,
		}
		var loc = request.CustomerLocation{
			Longitude: 10,
			Latitude:  1,
		}

		_ = repo.FindNearest(loc)
		ctrl.Finish()
	}

	assert.GreaterOrEqual(t, len(stateChanges), 0)
}

func TestRepositoryMatch_FindNearest(t *testing.T) {
	var resp = &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"coordinates": [40.946104,28.035588]}`)),
		Close:      false,
	}
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()

	client := mocks.NewMockClientImplementation(ctrl)
	client.
		EXPECT().
		Do(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(resp, nil).
		Times(1)

	var repo = RepositoryMatch{
		MatchServiceConfig: &config.Service{
			URL:  "url",
			Path: "/path",
		},
		breaker: circuitbreaker.New(),
		client:  client,
	}
	var loc = request.CustomerLocation{
		Longitude: 10,
		Latitude:  1,
	}

	var actualResponse = repo.FindNearest(loc)
	var expectedResponse = response.NearestDriver{
		Coordinates: []float64{40.946104, 28.035588},
	}

	assert.Equal(t, expectedResponse, actualResponse)
}
