package handler

import (
	"bit-driver-matching-service/domain/match"
	"bit-driver-matching-service/request"
	"bit-driver-matching-service/response"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	ErrorOnBind       = `{"message": "%s"}`
)

type MatchHandler struct {
	Service *match.Service
}

type MatchImplementations interface {
	FindNearestDriver(loc request.CustomerLocation) response.NearestDriver
}

func (h *MatchHandler) FindDriver(c echo.Context) error {
	var err error
	var cLocation request.CustomerLocation
	if err = c.Bind(&cLocation); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf(ErrorOnBind, err.Error()))
	}

	if err = c.Validate(&cLocation); err != nil {
		return err
	}

	var driver = h.Service.FindNearestDriver(cLocation)

	return c.JSON(http.StatusOK, driver)
}
