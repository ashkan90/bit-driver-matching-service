package main

import (
	"bit-driver-matching-service/adapters/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer() *echo.Echo {
	var sv = echo.New()
	sv.Validator = validator.NewRequestValidator()
	sv.Use(middleware.Recover())
	sv.Use(middleware.Logger())

	return sv
}