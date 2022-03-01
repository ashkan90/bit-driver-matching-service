package handler

import (
	"bit-driver-matching-service/adapters/validator"
	mocks "bit-driver-matching-service/domain/.mocks"
	"bit-driver-matching-service/domain/match"
	"bit-driver-matching-service/request"
	"bit-driver-matching-service/response"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMatchHandler_FindDriver(t *testing.T) {
	var loc = request.CustomerLocation{
		Longitude: 100,
		Latitude:  75,
	}
	var body = fmt.Sprintf(`{"longitude": %f, "latitude": %f}`, loc.Longitude, loc.Latitude)

	req := httptest.NewRequest(http.MethodGet, "/find-nearest", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	e.Validator = validator.NewRequestValidator()
	ctx := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var repoResponse = response.NearestDriver{
		Coordinates: []float64{100, 86.2},
	}
	var expectedResponse = `{"coordinates":[100,86.2]}` + "\n"

	mockRepo := mocks.NewMockRepositoryImplementation(ctrl)
	mockRepo.
		EXPECT().
		FindNearest(loc).
		Return(repoResponse).
		Times(1)

	handler := &MatchHandler{
		Service: match.NewService(mockRepo),
	}

	err := handler.FindDriver(ctx)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, rec.Body.String())
}
