package match

import (
	"bit-driver-matching-service/request"
	"bit-driver-matching-service/response"
)

type RepositoryMatch struct {}

func NewRepositoryMatch() *RepositoryMatch {
	return &RepositoryMatch{}
}

func (r *RepositoryMatch) FindNearest(loc request.CustomerLocation) response.NearestDriver {

	return response.NearestDriver{
		Name:        "Emirhan",
		Coordinates: []float64{41.08712, 29.156156},
	}
}

