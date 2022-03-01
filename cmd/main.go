package main

import (
	"bit-driver-matching-service/adapters/rest"
	"bit-driver-matching-service/config"
	"context"
	"fmt"
	"github.com/benbjohnson/clock"
	"github.com/cenkalti/backoff/v3"
	"github.com/labstack/echo/v4"
	"github.com/mercari/go-circuitbreaker"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Rest adapter which serves the routes via given settings
	var port = os.Getenv("PORT")
	var conf = config.NewGeneralConfig("service_config.yaml")
	var logger = log.Default()
	var sv = NewServer()

	if port != "" {
		conf.Server.Port = port
	}

	var cb = circuitbreaker.New(
		circuitbreaker.WithClock(clock.New()),
		circuitbreaker.WithFailOnContextCancel(true),
		circuitbreaker.WithFailOnContextDeadline(true),
		circuitbreaker.WithHalfOpenMaxSuccesses(10),
		circuitbreaker.WithOpenTimeoutBackOff(backoff.NewExponentialBackOff()),
		circuitbreaker.WithOpenTimeout(10*time.Second),
		circuitbreaker.WithCounterResetInterval(10*time.Second),
		circuitbreaker.WithTripFunc(circuitbreaker.NewTripFuncFailureRate(10, 0.4)),
		circuitbreaker.WithOnStateChangeHookFn(func(from, to circuitbreaker.State) {
			logger.Printf("state changed from %s to %s\n", from, to)
		}),
	)

	var restAdapter = &rest.Adapter{
		Config:  conf.Server,
		Logger:  logger,
		Server:  sv,
		Breaker: cb,
	}
	restAdapter.Serve()

	go func() {
		logger.Fatal(sv.Start(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)))
	}()

	gracefulShutdown(logger, sv)
}

func gracefulShutdown(logger *log.Logger, s *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		logger.Fatal(err)
	}
}
