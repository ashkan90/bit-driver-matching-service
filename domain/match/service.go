package match

import (
	"bit-driver-matching-service/request"
	"bit-driver-matching-service/response"
)

type Service struct {
	Repository *RepositoryMatch
}

type RepositoryImplementation interface {
	FindNearest(loc request.CustomerLocation) response.NearestDriver
}

func NewService(repo *RepositoryMatch) *Service {
	return &Service{Repository: repo}
}

func (s *Service) FindNearestDriver(loc request.CustomerLocation) response.NearestDriver {
	return s.Repository.FindNearest(loc)
}

