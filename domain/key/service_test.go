package key

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	var fileRepo, _ = NewFileRepository("/")
	var service = NewService(fileRepo, 0)

	assert.NotNil(t, service)
	assert.NotNil(t, service.keys)
}

func TestService_Set(t *testing.T) {
	var fileRepo, _ = NewFileRepository("/")
	var service = NewService(fileRepo, 0)
	var _error = service.Set("myKey", "myVal")

	var expectedValue = Keys{
		"myKey": "myVal",
	}

	assert.Nil(t, _error)
	assert.Equal(t, expectedValue, service.keys)
}

func TestService_SetFailOnEmptyKeyString(t *testing.T) {
	var fileRepo, _ = NewFileRepository("/")
	var service = NewService(fileRepo, 0)
	var _error = service.Set("", "myVal")

	assert.NotNil(t, _error)
}

func TestService_Get(t *testing.T) {
	// mock the repository

}
