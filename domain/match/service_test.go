package match

import (
	mocks "bit-driver-matching-service/domain/.mocks"
	"bit-driver-matching-service/request"
	"bit-driver-matching-service/response"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_FindNearestDriverLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var repoResponse = response.NearestDriver{
		Coordinates: []float64{110, 75.21},
	}
	var loc = request.CustomerLocation{
		Longitude: 100,
		Latitude:  75,
	}

	mockRepo := mocks.NewMockRepositoryImplementation(ctrl)
	mockRepo.
		EXPECT().
		FindNearest(loc).
		Return(repoResponse).
		Times(1)

	service := NewService(mockRepo)

	var actualResponse = service.FindNearestDriver(loc)
	var expectedResponse = []float64{110, 75.21}

	assert.Equal(t, expectedResponse, actualResponse.Coordinates)
}
