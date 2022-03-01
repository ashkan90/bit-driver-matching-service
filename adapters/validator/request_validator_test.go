package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type requestObj struct {
	Name    string `json:"name" validate:"required"`
	Surname string `json:"surname" validate:"required"`
	Age     int    `json:"age" validate:"required"`
}

func TestRequestValidator_ValidateShouldReturnError(t *testing.T) {
	var validator = NewRequestValidator()
	var obj requestObj

	err := validator.Validate(&obj)

	assert.NotNil(t, err)
}

func TestRequestValidator_ValidateShouldNotReturnError(t *testing.T) {
	var validator = NewRequestValidator()
	var obj = requestObj{
		Name:    "my-name",
		Surname: "surname",
		Age:     10,
	}

	err := validator.Validate(&obj)

	assert.Nil(t, err)
}
