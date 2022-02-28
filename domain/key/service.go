package key

import (
	"errors"
	"log"
	"sync"
	"time"
)

const (
	ErrorOnGetKeyNotFound = "key not found on db"
)

type Keys map[string]string

type Service struct {
	interval   int
	keys       Keys
	Repository *FileRepository

	mu sync.Mutex
}

func NewService(repo *FileRepository, interval int) *Service {
	var service = &Service{
		Repository: repo,
		interval:   interval,
		mu:         sync.Mutex{},
	}

	service.keys = service.getPersist()

	go service.setPersist()

	return service
}

func (s *Service) Set(k, v string) error {
	if k == "" {
		return errors.New("key cannot be empty")
	}

	s.mu.Lock()
	s.keys[k] = v
	s.mu.Unlock()

	return nil
}

func (s *Service) Get(k string) (string, error) {
	if val, ok := s.keys[k]; ok {
		return val, nil
	}
	return "", errors.New(ErrorOnGetKeyNotFound)
}

// nolint:gosimple
func (s *Service) setPersist() {
	if s.interval <= 0 {
		return
	}

	var ticker = time.NewTicker(time.Duration(s.interval) * time.Second)

	for {
		select {
		case <-ticker.C:
			s.mu.Lock()
			err := s.Repository.WriteToFile(s.keys)
			if err != nil {
				log.Println("Something went wrong while doing persistent writing process. Details: " + err.Error())
			}
			s.mu.Unlock()
		}
	}
}

func (s *Service) getPersist() Keys {
	var keys, err = s.Repository.ReadFromFile()
	if err != nil {
		return Keys{}
	}

	return keys
}
