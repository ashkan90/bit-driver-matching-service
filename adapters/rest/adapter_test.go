package rest

import (
	"bit-driver-matching-service/config"
	"log"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAdapter_ServeShouldAddRespectiveRoutesCorrectly(t *testing.T) {
	var adapter = &Adapter{
		Config: config.Server{
			Port: "1111",
			Service: config.Service{
				URL: "http://other-service.com/",
			},
		},
		Logger: log.Default(),
		Server: echo.New(),
	}

	adapter.Serve()

	assert.Len(t, adapter.Server.Routes(), 1)
}
