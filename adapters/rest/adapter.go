package rest

import (
	"bit-driver-matching-service/adapters/handler"
	"bit-driver-matching-service/config"
	"bit-driver-matching-service/domain/match"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Adapter struct {
	Config config.Server
	Logger *log.Logger
	Server *echo.Echo
}

func (a *Adapter) Serve() {
	var matchRepository = match.NewRepositoryMatch(&a.Config.Service)
	var matchService = match.NewService(matchRepository)
	var matchRest = &handler.MatchHandler{Service: matchService}

	a.Server.Add(http.MethodGet, "/find-nearest", matchRest.FindDriver)

	a.Logger.Println("Server will be started on port " + a.Config.Port)
}
