package rest

import (
	"bit-driver-matching-service/adapters/handler"
	"bit-driver-matching-service/config"
	"bit-driver-matching-service/domain/match"
	"github.com/labstack/echo/v4"
	"github.com/mercari/go-circuitbreaker"
	"log"
	"net/http"
)

type Adapter struct {
	Config  config.Server
	Logger  *log.Logger
	Server  *echo.Echo
	Breaker *circuitbreaker.CircuitBreaker
}

func (a *Adapter) Serve() {
	var matchRepository = match.NewRepositoryMatch(&a.Config.Service, a.Breaker)
	var matchService = match.NewService(matchRepository)
	var matchRest = &handler.MatchHandler{Service: matchService}

	a.Server.Add(http.MethodGet, "/find-nearest", matchRest.FindDriver)

	a.Logger.Println("Server will be started on port " + a.Config.Port)
}
